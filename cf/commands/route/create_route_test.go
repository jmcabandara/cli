package route_test

import (
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/commands/route"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/simonleung8/flags"

	fakeapi "github.com/cloudfoundry/cli/cf/api/fakes"
	fakerequirements "github.com/cloudfoundry/cli/cf/requirements/fakes"

	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/cloudfoundry/cli/testhelpers/matchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateRoute", func() {
	var (
		ui         *testterm.FakeUI
		routeRepo  *fakeapi.FakeRouteRepository
		configRepo core_config.Repository

		cmd         command_registry.Command
		deps        command_registry.Dependency
		factory     *fakerequirements.FakeFactory
		flagContext flags.FlagContext

		spaceRequirement  requirements.SpaceRequirement
		domainRequirement requirements.DomainRequirement
		// minAPIVersionRequirement requirements.Requirement
	)

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		configRepo = testconfig.NewRepositoryWithDefaults()
		routeRepo = &fakeapi.FakeRouteRepository{}
		repoLocator := deps.RepoLocator.SetRouteRepository(routeRepo)

		deps = command_registry.Dependency{
			Ui:          ui,
			Config:      configRepo,
			RepoLocator: repoLocator,
		}

		cmd = &route.CreateRoute{}
		cmd.SetDependency(deps, false)

		flagContext = flags.NewFlagContext(cmd.MetaData().Flags)

		factory = &fakerequirements.FakeFactory{}

		spaceRequirement = &fakerequirements.FakeSpaceRequirement{}
		factory.NewSpaceRequirementReturns(spaceRequirement)

		domainRequirement = &fakerequirements.FakeDomainRequirement{}
		factory.NewDomainRequirementReturns(domainRequirement)

		// minAPIVersionRequirement = &passingRequirement{}
		// factory.NewMinAPIVersionRequirementReturns(minAPIVersionRequirement)
	})

	Describe("Requirements", func() {
		Context("when not provided exactly two args", func() {
			BeforeEach(func() {
				flagContext.Parse("space-name")
			})

			It("fails with usage", func() {
				Expect(func() { cmd.Requirements(factory, flagContext) }).To(Panic())
				Expect(ui.Outputs).To(ContainSubstrings(
					[]string{"Incorrect Usage. Requires SPACE and DOMAIN as arguments"},
					[]string{"NAME"},
					[]string{"USAGE"},
				))
			})
		})

		Context("when provided exactly two args", func() {
			BeforeEach(func() {
				flagContext.Parse("space-name", "domain-name")
			})

			It("returns a SpaceRequirement", func() {
				actualRequirements, err := cmd.Requirements(factory, flagContext)
				Expect(err).NotTo(HaveOccurred())
				Expect(factory.NewSpaceRequirementCallCount()).To(Equal(1))
				Expect(factory.NewSpaceRequirementArgsForCall(0)).To(Equal("space-name"))

				Expect(actualRequirements).To(ContainElement(spaceRequirement))
			})

			It("returns a DomainRequirement", func() {
				actualRequirements, err := cmd.Requirements(factory, flagContext)
				Expect(err).NotTo(HaveOccurred())
				Expect(factory.NewDomainRequirementCallCount()).To(Equal(1))
				Expect(factory.NewDomainRequirementArgsForCall(0)).To(Equal("domain-name"))

				Expect(actualRequirements).To(ContainElement(domainRequirement))
			})
		})
	})
})

// var _ = Describe("create-route command", func() {
// 	var (
// 		ui                  *testterm.FakeUI
// 		routeRepo           *testapi.FakeRouteRepository
// 		requirementsFactory *testreq.FakeReqFactory
// 		config              core_config.Repository
// 		deps                command_registry.Dependency
// 	)

// 	updateCommandDependency := func(pluginCall bool) {
// 		deps.Ui = ui
// 		deps.RepoLocator = deps.RepoLocator.SetRouteRepository(routeRepo)
// 		deps.Config = config
// 		command_registry.Commands.SetCommand(command_registry.Commands.FindCommand("create-route").SetDependency(deps, pluginCall))
// 	}

// 	BeforeEach(func() {
// 		ui = &testterm.FakeUI{}
// 		routeRepo = &testapi.FakeRouteRepository{}
// 		requirementsFactory = &testreq.FakeReqFactory{}
// 		config = testconfig.NewRepositoryWithDefaults()
// 	})

// 	runCommand := func(args ...string) bool {
// 		return testcmd.RunCliCommand("create-route", args, requirementsFactory, updateCommandDependency, false)
// 	}

// 	Describe("requirements", func() {
// 		It("fails when not logged in", func() {
// 			requirementsFactory.TargetedOrgSuccess = true

// 			Expect(runCommand("my-space", "example.com", "-n", "foo")).To(BeFalse())
// 		})

// 		It("fails when an org is not targeted", func() {
// 			requirementsFactory.LoginSuccess = true

// 			Expect(runCommand("my-space", "example.com", "-n", "foo")).To(BeFalse())
// 		})

// 		It("fails with usage when not provided two args", func() {
// 			requirementsFactory.LoginSuccess = true
// 			requirementsFactory.TargetedOrgSuccess = true

// 			runCommand("my-space")
// 			Expect(ui.Outputs).To(ContainSubstrings(
// 				[]string{"Incorrect Usage", "Requires", "arguments"},
// 				[]string{"create-route SPACE DOMAIN [-n HOSTNAME] [--path PATH]"},
// 			))
// 		})
// 	})

// 	Context("when logged in, targeted a space and given a domain that exists", func() {
// 		BeforeEach(func() {
// 			requirementsFactory.LoginSuccess = true
// 			requirementsFactory.TargetedOrgSuccess = true
// 			requirementsFactory.Domain = models.DomainFields{
// 				Guid: "domain-guid",
// 				Name: "example.com",
// 			}
// 			requirementsFactory.Space = models.Space{SpaceFields: models.SpaceFields{
// 				Guid: "my-space-guid",
// 				Name: "my-space",
// 			}}
// 		})

// 		It("creates routes, obviously", func() {
// 			runCommand("-n", "host", "my-space", "example.com")

// 			Expect(ui.Outputs).To(ContainSubstrings(
// 				[]string{"Creating route", "host.example.com", "my-org", "my-space", "my-user"},
// 				[]string{"OK"},
// 			))

// 			Expect(routeRepo.CreateInSpaceHost).To(Equal("host"))
// 			Expect(routeRepo.CreateInSpaceDomainGuid).To(Equal("domain-guid"))
// 			Expect(routeRepo.CreateInSpaceSpaceGuid).To(Equal("my-space-guid"))
// 		})

// 		It("creates routes with a context path", func() {
// 			runCommand("-n", "host", "--path", "path", "my-space", "example.com")

// 			Expect(ui.Outputs).To(ContainSubstrings(
// 				[]string{"Creating route", "host.example.com/path", "my-org", "my-space", "my-user"},
// 				[]string{"OK"},
// 			))

// 			Expect(routeRepo.CreateInSpaceHost).To(Equal("host"))
// 			Expect(routeRepo.CreateInSpacePath).To(Equal("/path"))
// 			Expect(routeRepo.CreateInSpaceDomainGuid).To(Equal("domain-guid"))
// 			Expect(routeRepo.CreateInSpaceSpaceGuid).To(Equal("my-space-guid"))
// 		})

// 		It("is idempotent", func() {
// 			routeRepo.CreateInSpaceErr = true
// 			routeRepo.FindByHostAndDomainReturns.Route = models.Route{
// 				Space:  requirementsFactory.Space.SpaceFields,
// 				Guid:   "my-route-guid",
// 				Host:   "host",
// 				Path:   "/path",
// 				Domain: requirementsFactory.Domain,
// 			}

// 			runCommand("-n", "host", "my-space", "example.com")

// 			Expect(ui.Outputs).To(ContainSubstrings(
// 				[]string{"Creating route"},
// 				[]string{"OK"},
// 				[]string{"host.example.com/path", "already exists"},
// 			))

// 			Expect(routeRepo.CreateInSpaceHost).To(Equal("host"))
// 			Expect(routeRepo.CreateInSpaceDomainGuid).To(Equal("domain-guid"))
// 			Expect(routeRepo.CreateInSpaceSpaceGuid).To(Equal("my-space-guid"))
// 		})

// 		Describe("RouteCreator interface", func() {
// 			It("creates a route", func() {
// 				createdRoute := models.Route{}
// 				createdRoute.Host = "my-host"
// 				createdRoute.Guid = "my-route-guid"
// 				routeRepo = &testapi.FakeRouteRepository{
// 					CreateInSpaceCreatedRoute: createdRoute,
// 				}

// 				updateCommandDependency(false)
// 				c := command_registry.Commands.FindCommand("create-route")
// 				cmd := c.(RouteCreator)
// 				route, apiErr := cmd.CreateRoute("my-host", "/path", requirementsFactory.Domain, requirementsFactory.Space.SpaceFields)

// 				Expect(apiErr).NotTo(HaveOccurred())
// 				Expect(route.Guid).To(Equal(createdRoute.Guid))
// 				Expect(ui.Outputs).To(ContainSubstrings(
// 					[]string{"Creating route", "my-host.example.com/path", "my-org", "my-space", "my-user"},
// 					[]string{"OK"},
// 				))

// 				Expect(routeRepo.CreateInSpaceHost).To(Equal("my-host"))
// 				Expect(routeRepo.CreateInSpaceDomainGuid).To(Equal("domain-guid"))
// 				Expect(routeRepo.CreateInSpaceSpaceGuid).To(Equal("my-space-guid"))
// 				Expect(routeRepo.CreateInSpacePath).To(Equal("/path"))
// 			})
// 		})
// 	})
// })
