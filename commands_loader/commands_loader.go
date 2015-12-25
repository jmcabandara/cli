package commands_loader

import (
	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/commands"
	"github.com/cloudfoundry/cli/cf/commands/application"
	"github.com/cloudfoundry/cli/cf/commands/buildpack"
	"github.com/cloudfoundry/cli/cf/commands/domain"
	"github.com/cloudfoundry/cli/cf/commands/environmentvariablegroup"
	"github.com/cloudfoundry/cli/cf/commands/featureflag"
	"github.com/cloudfoundry/cli/cf/commands/organization"
	"github.com/cloudfoundry/cli/cf/commands/plugin"
	"github.com/cloudfoundry/cli/cf/commands/plugin_repo"
	"github.com/cloudfoundry/cli/cf/commands/quota"
	"github.com/cloudfoundry/cli/cf/commands/route"
	"github.com/cloudfoundry/cli/cf/commands/routergroups"
	"github.com/cloudfoundry/cli/cf/commands/securitygroup"
	"github.com/cloudfoundry/cli/cf/commands/service"
	"github.com/cloudfoundry/cli/cf/commands/serviceaccess"
	"github.com/cloudfoundry/cli/cf/commands/serviceauthtoken"
	"github.com/cloudfoundry/cli/cf/commands/servicebroker"
	"github.com/cloudfoundry/cli/cf/commands/servicekey"
	"github.com/cloudfoundry/cli/cf/commands/space"
	"github.com/cloudfoundry/cli/cf/commands/spacequota"
	"github.com/cloudfoundry/cli/cf/commands/user"
)

func Load() {
	command_registry.Register(&application.ShowApp{})            // app
	command_registry.Register(&application.ListApps{})           // apps
	command_registry.Register(&application.CopySource{})         // copy-source
	command_registry.Register(&application.DeleteApp{})          // delete
	command_registry.Register(&application.DisableSSH{})         // disable-ssh
	command_registry.Register(&application.EnableSSH{})          // enable-ssh
	command_registry.Register(&application.Env{})                // env
	command_registry.Register(&application.Events{})             // events
	command_registry.Register(&application.Files{})              // files
	command_registry.Register(&application.GetHealthCheck{})     // get-health-check
	command_registry.Register(&application.Logs{})               // logs
	command_registry.Register(&application.Push{})               // push
	command_registry.Register(&application.RenameApp{})          // rename
	command_registry.Register(&application.Restage{})            // restage
	command_registry.Register(&application.Restart{})            // restart
	command_registry.Register(&application.RestartAppInstance{}) // restart-app-instance
	command_registry.Register(&application.Scale{})              // scale
	command_registry.Register(&application.SetEnv{})             // set-env
	command_registry.Register(&application.SetHealthCheck{})     // set-health-check
	command_registry.Register(&application.SSH{})                // ssh
	command_registry.Register(&application.SSHEnabled{})         // ssh-enabled
	command_registry.Register(&application.Start{})              // start
	command_registry.Register(&application.Stop{})               // stop
	command_registry.Register(&application.UnsetEnv{})           // unset-env

	command_registry.Register(&commands.Api{})               // api
	command_registry.Register(&commands.Authenticate{})      // auth
	command_registry.Register(&commands.ConfigCommands{})    // config
	command_registry.Register(&commands.CreateAppManifest{}) // create-app-manifest
	command_registry.Register(&commands.Curl{})              // curl
	command_registry.Register(&commands.Help{})              // help
	command_registry.Register(&commands.Login{})             // login
	command_registry.Register(&commands.Logout{})            // logout
	command_registry.Register(&commands.OAuthToken{})        // oauth-token
	command_registry.Register(&commands.Password{})          // passwd
	command_registry.Register(&commands.OneTimeSSHCode{})    // ssh-code
	command_registry.Register(&commands.ListStack{})         // stack
	command_registry.Register(&commands.ListStacks{})        // stacks
	command_registry.Register(&commands.Target{})            // target

	command_registry.Register(&buildpack.ListBuildpacks{})  // buildpacks
	command_registry.Register(&buildpack.CreateBuildpack{}) // create-buildpack
	command_registry.Register(&buildpack.DeleteBuildpack{}) // delete-buildpack
	command_registry.Register(&buildpack.RenameBuildpack{}) // rename-buildpack
	command_registry.Register(&buildpack.UpdateBuildpack{}) // update-buildpack

	command_registry.Register(&domain.CreateDomain{})       // create-domain
	command_registry.Register(&domain.CreateSharedDomain{}) // create-shared-domain
	command_registry.Register(&domain.DeleteDomain{})       // delete-domain
	command_registry.Register(&domain.DeleteSharedDomain{}) // delete-shared-domain
	command_registry.Register(&domain.ListDomains{})        // domains

	command_registry.Register(&environmentvariablegroup.RunningEnvironmentVariableGroup{})    // running-environment-variable-group
	command_registry.Register(&environmentvariablegroup.SetRunningEnvironmentVariableGroup{}) // set-running-environment-variable-group
	command_registry.Register(&environmentvariablegroup.SetStagingEnvironmentVariableGroup{}) // set-staging-environment-variable-group
	command_registry.Register(&environmentvariablegroup.StagingEnvironmentVariableGroup{})    // staging-environment-variable-group

	command_registry.Register(&featureflag.DisableFeatureFlag{}) // disable-feature-flag
	command_registry.Register(&featureflag.EnableFeatureFlag{})  // enable-feature-flag
	command_registry.Register(&featureflag.ShowFeatureFlag{})    // feature-flag
	command_registry.Register(&featureflag.ListFeatureFlags{})   // feature-flags

	command_registry.Register(&organization.CreateOrg{})            // create-org
	command_registry.Register(&organization.DeleteOrg{})            // delete-org
	command_registry.Register(&organization.ShowOrg{})              // org
	command_registry.Register(&organization.ListOrgs{})             // orgs
	command_registry.Register(&organization.RenameOrg{})            // rename-org
	command_registry.Register(&organization.SetQuota{})             // set-quota
	command_registry.Register(&organization.SharePrivateDomain{})   // share-private-domain
	command_registry.Register(&organization.UnsharePrivateDomain{}) // unshare-private-domain

	command_registry.Register(&plugin.PluginInstall{})   // install-plugin
	command_registry.Register(&plugin.Plugins{})         // plugins
	command_registry.Register(&plugin.PluginUninstall{}) // uninstall-plugin

	command_registry.Register(&plugin_repo.AddPluginRepo{})    // add-plugin-repo
	command_registry.Register(&plugin_repo.ListPluginRepos{})  // list-plugin-repos
	command_registry.Register(&plugin_repo.RemovePluginRepo{}) // remove-plugin-repo
	command_registry.Register(&plugin_repo.RepoPlugins{})      // repo-plugins

	command_registry.Register(&quota.CreateQuota{}) // create-quota
	command_registry.Register(&quota.DeleteQuota{}) // delete-quota
	command_registry.Register(&quota.ShowQuota{})   // quota

	command_registry.Register(&route.CheckRoute{})           // check-route
	command_registry.Register(&route.CreateRoute{})          // create-route
	command_registry.Register(&route.DeleteOrphanedRoutes{}) // delete-orphaned-routes
	command_registry.Register(&route.DeleteRoute{})          // delete-route
	command_registry.Register(&route.MapRoute{})             // map-route
	command_registry.Register(&route.ListRoutes{})           // routes
	command_registry.Register(&route.UnmapRoute{})           // unmap-route

	command_registry.Register(&routergroups.RouterGroups{}) // router-groups

	command_registry.Register(&securitygroup.BindToRunningGroup{})        // bind-running-security-group
	command_registry.Register(&securitygroup.BindSecurityGroup{})         // bind-security-group
	command_registry.Register(&securitygroup.BindToStagingGroup{})        // bind-staging-security-group
	command_registry.Register(&securitygroup.CreateSecurityGroup{})       // create-security-group
	command_registry.Register(&securitygroup.DeleteSecurityGroup{})       // delete-security-group
	command_registry.Register(&securitygroup.ListRunningSecurityGroups{}) // running-security-groups
	command_registry.Register(&securitygroup.ShowSecurityGroup{})         // security-group
	command_registry.Register(&securitygroup.SecurityGroups{})            // security-groups
	command_registry.Register(&securitygroup.ListStagingSecurityGroups{}) // staging-security-groups
	command_registry.Register(&securitygroup.UnbindFromRunningGroup{})    // unbind-running-security-group
	command_registry.Register(&securitygroup.UnbindSecurityGroup{})       // unbind-security-group
	command_registry.Register(&securitygroup.UnbindFromStagingGroup{})    // unbind-staging-security-group
	command_registry.Register(&securitygroup.UpdateSecurityGroup{})       // update-security-group

	command_registry.Register(&service.BindService{})               // bind-service
	command_registry.Register(&service.CreateService{})             // create-service
	command_registry.Register(&service.CreateUserProvidedService{}) // create-user-provided-service
	command_registry.Register(&service.DeleteService{})             // delete-service
	command_registry.Register(&service.MarketplaceServices{})       // marketplace
	command_registry.Register(&service.MigrateServiceInstances{})   // migrate-service-instances
	command_registry.Register(&service.PurgeServiceInstance{})      // purge-service-instance
	command_registry.Register(&service.PurgeServiceOffering{})      // purge-service-offering
	command_registry.Register(&service.RenameService{})             // rename-service
	command_registry.Register(&service.ShowService{})               // service
	command_registry.Register(&service.ListServices{})              // services
	command_registry.Register(&service.UnbindService{})             // unbind-service
	command_registry.Register(&service.UpdateService{})             // update-service
	command_registry.Register(&service.UpdateUserProvidedService{}) // update-user-provided-service

	command_registry.Register(&serviceaccess.DisableServiceAccess{}) // disable-service-access
	command_registry.Register(&serviceaccess.EnableServiceAccess{})  // enable-service-access
	command_registry.Register(&serviceaccess.ServiceAccess{})        // service-access

	command_registry.Register(&serviceauthtoken.CreateServiceAuthTokenFields{}) // create-service-auth-token
	command_registry.Register(&serviceauthtoken.DeleteServiceAuthTokenFields{}) // delete-service-auth-token
	command_registry.Register(&serviceauthtoken.ListServiceAuthTokens{})        // service-auth-tokens
	command_registry.Register(&serviceauthtoken.UpdateServiceAuthTokenFields{}) // update-service-auth-token

	command_registry.Register(&servicebroker.CreateServiceBroker{}) // create-service-broker
	command_registry.Register(&servicebroker.DeleteServiceBroker{}) // delete-service-broker
	command_registry.Register(&servicebroker.RenameServiceBroker{}) // rename-service-broker
	command_registry.Register(&servicebroker.ListServiceBrokers{})  // service-brokers
	command_registry.Register(&servicebroker.UpdateServiceBroker{}) // update-service-broker

	command_registry.Register(&servicekey.CreateServiceKey{}) // create-service-key
	command_registry.Register(&servicekey.DeleteServiceKey{}) // delete-service-key
	command_registry.Register(&servicekey.ServiceKey{})       // service-key
	command_registry.Register(&servicekey.ServiceKeys{})      // service-keys

	command_registry.Register(&space.AllowSpaceSSH{})    // allow-space-ssh
	command_registry.Register(&space.CreateSpace{})      // create-space
	command_registry.Register(&space.DeleteSpace{})      // delete-space
	command_registry.Register(&space.DisallowSpaceSSH{}) // disallow-space-ssh
	command_registry.Register(&space.RenameSpace{})      // rename-space
	command_registry.Register(&space.ShowSpace{})        // space
	command_registry.Register(&space.SpaceSSHAllowed{})  // space-ssh-allowed
	command_registry.Register(&space.ListSpaces{})       // spaces

	command_registry.Register(&spacequota.CreateSpaceQuota{}) // create-space-quota
	command_registry.Register(&spacequota.DeleteSpaceQuota{}) // delete-space-quota
	command_registry.Register(&spacequota.SetSpaceQuota{})    // set-space-quota
	command_registry.Register(&spacequota.SpaceQuota{})       // space-quota
	command_registry.Register(&spacequota.ListSpaceQuotas{})  // space-quotas
	command_registry.Register(&spacequota.UnsetSpaceQuota{})  // unset-space-quota
	command_registry.Register(&spacequota.UpdateSpaceQuota{}) // update-space-quota

	command_registry.Register(&user.CreateUser{})     // create-user
	command_registry.Register(&user.DeleteUser{})     // delete-user
	command_registry.Register(&user.OrgUsers{})       // org-users
	command_registry.Register(&user.SetOrgRole{})     // set-org-role
	command_registry.Register(&user.SetSpaceRole{})   // set-space-role
	command_registry.Register(&user.SpaceUsers{})     // space-users
	command_registry.Register(&user.UnsetOrgRole{})   // unset-org-role
	command_registry.Register(&user.UnsetSpaceRole{}) // unset-space-role
}
