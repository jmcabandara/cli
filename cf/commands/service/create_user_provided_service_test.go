package service_test

import (
	"fmt"
	"io/ioutil"
	"os"

	testapi "github.com/cloudfoundry/cli/cf/api/fakes"
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testreq "github.com/cloudfoundry/cli/testhelpers/requirements"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/cli/testhelpers/matchers"
)

var _ = Describe("create-user-provided-service command", func() {
	var (
		ui                  *testterm.FakeUI
		config              core_config.Repository
		repo                *testapi.FakeUserProvidedServiceInstanceRepository
		requirementsFactory *testreq.FakeReqFactory
		deps                command_registry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.Ui = ui
		deps.RepoLocator = deps.RepoLocator.SetUserProvidedServiceInstanceRepository(repo)
		deps.Config = config
		command_registry.Commands.SetCommand(command_registry.Commands.FindCommand("create-user-provided-service").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		config = testconfig.NewRepositoryWithDefaults()
		repo = &testapi.FakeUserProvidedServiceInstanceRepository{}
		requirementsFactory = &testreq.FakeReqFactory{LoginSuccess: true, TargetedSpaceSuccess: true}
	})

	Describe("login requirements", func() {
		It("fails if the user is not logged in", func() {
			requirementsFactory.LoginSuccess = false
			Expect(testcmd.RunCliCommand("create-user-provided-service", []string{"my-service"}, requirementsFactory, updateCommandDependency, false)).To(BeFalse())
		})
		It("fails when a space is not targeted", func() {
			requirementsFactory.TargetedSpaceSuccess = false
			Expect(testcmd.RunCliCommand("create-user-provided-service", []string{"my-service"}, requirementsFactory, updateCommandDependency, false)).To(BeFalse())
		})
	})

	It("creates a new user provided service given just a name", func() {
		testcmd.RunCliCommand("create-user-provided-service", []string{"my-custom-service"}, requirementsFactory, updateCommandDependency, false)
		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"Creating user provided service"},
			[]string{"OK"},
		))
	})

	It("accepts service parameters interactively", func() {
		ui.Inputs = []string{"foo value", "bar value", "baz value"}
		testcmd.RunCliCommand("create-user-provided-service", []string{"-p", `"foo, bar, baz"`, "my-custom-service"}, requirementsFactory, updateCommandDependency, false)

		Expect(ui.Prompts).To(ContainSubstrings(
			[]string{"foo"},
			[]string{"bar"},
			[]string{"baz"},
		))

		Expect(repo.CreateCallCount()).To(Equal(1))
		name, drainUrl, _, params := repo.CreateArgsForCall(0)
		Expect(name).To(Equal("my-custom-service"))
		Expect(drainUrl).To(Equal(""))
		Expect(params["foo"]).To(Equal("foo value"))
		Expect(params["bar"]).To(Equal("bar value"))
		Expect(params["baz"]).To(Equal("baz value"))

		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"Creating user provided service", "my-custom-service", "my-org", "my-space", "my-user"},
			[]string{"OK"},
		))
	})

	It("accepts service parameters as single-quoted JSON without prompting", func() {
		args := []string{"-p", `'{"foo": "foo value", "bar": "bar value", "baz": 4}'`, "my-custom-service"}
		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)

		name, _, _, params := repo.CreateArgsForCall(0)
		Expect(name).To(Equal("my-custom-service"))

		Expect(ui.Prompts).To(BeEmpty())
		Expect(params).To(Equal(map[string]interface{}{
			"foo": "foo value",
			"bar": "bar value",
			"baz": float64(4),
		}))

		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"Creating user provided service"},
			[]string{"OK"},
		))
	})

	It("accepts service parameters as double-quoted JSON without prompting", func() {
		args := []string{"-p", `"{"foo": "foo value", "bar": "bar value", "baz": 4}"`, "my-custom-service"}
		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)

		name, _, _, params := repo.CreateArgsForCall(0)
		Expect(name).To(Equal("my-custom-service"))

		Expect(ui.Prompts).To(BeEmpty())
		Expect(params).To(Equal(map[string]interface{}{
			"foo": "foo value",
			"bar": "bar value",
			"baz": float64(4),
		}))

		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"Creating user provided service"},
			[]string{"OK"},
		))
	})

	It("fails with an error when given bad literal JSON", func() {
		args := []string{"-p", `'{:bad_json:}'`, "my-custom-service"}
		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)
		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"FAILED"},
		))
	})

	It("accepts service parameters as a file containing JSON without prompting", func() {
		tempfile, err := ioutil.TempFile("", "create-user-provided-service-test")
		Expect(err).NotTo(HaveOccurred())
		jsonData := `{"foo": "foo value", "bar": "bar value", "baz": 4}`
		ioutil.WriteFile(tempfile.Name(), []byte(jsonData), os.ModePerm)
		args := []string{"-p", fmt.Sprintf("@%s", tempfile.Name()), "my-custom-service"}

		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)

		name, _, _, params := repo.CreateArgsForCall(0)
		Expect(name).To(Equal("my-custom-service"))

		Expect(ui.Prompts).To(BeEmpty())
		Expect(params).To(Equal(map[string]interface{}{
			"foo": "foo value",
			"bar": "bar value",
			"baz": float64(4),
		}))

		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"Creating user provided service"},
			[]string{"OK"},
		))
	})

	It("fails with an error when given a file containing bad JSON", func() {
		tempfile, err := ioutil.TempFile("", "create-user-provided-service-test")
		Expect(err).NotTo(HaveOccurred())
		jsonData := `{:bad_json:}`
		ioutil.WriteFile(tempfile.Name(), []byte(jsonData), os.ModePerm)
		args := []string{"-p", fmt.Sprintf("@%s", tempfile.Name()), "my-custom-service"}

		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)
		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"FAILED"},
		))
	})

	It("fails with an error when given a file that cannot be read", func() {
		args := []string{"-p", "@nonexistent-file", "my-custom-service"}
		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)
		Expect(ui.Outputs).To(ContainSubstrings(
			[]string{"FAILED"},
		))
	})

	It("calls the create api with the corresponding syslog drain url", func() {
		args := []string{"-l", "syslog://example.com", "my-custom-service"}
		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)

		_, drainUrl, _, _ := repo.CreateArgsForCall(0)
		Expect(drainUrl).To(Equal("syslog://example.com"))
	})

	It("calls the create api with the corresponding route service url", func() {
		args := []string{"-r", "https://example.com", "my-custom-service"}
		testcmd.RunCliCommand("create-user-provided-service", args, requirementsFactory, updateCommandDependency, false)

		_, _, routeServiceUrl, _ := repo.CreateArgsForCall(0)
		Expect(routeServiceUrl).To(Equal("https://example.com"))
	})
})
