package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/flags"
)

const (
	iaasConfigurationPath              = "/infrastructure/iaas_configuration/edit"
	directorConfigurationPath          = "/infrastructure/director_configuration/edit"
	securityConfigurationPath          = "/infrastructure/security_tokens/edit"
	availabilityZonesConfigurationPath = "/infrastructure/availability_zones/edit"
	networksConfigurationPath          = "/infrastructure/networks/edit"
)

type ConfigureBosh struct {
	service boshFormService
	logger  logger
	Options struct {
		IaaSConfiguration              string `short:"i"  long:"iaas-configuration"  description:"iaas specific JSON configuration for the bosh director"`
		DirectorConfiguration          string `short:"d"  long:"director-configuration"  description:"director-specific JSON configuration for the bosh director"`
		SecurityConfiguration          string `short:"s"  long:"security-configuration"  description:"security-specific JSON configuration for the bosh director"`
		AvailabilityZonesConfiguration string `short:"a"  long:"az-configuration"  description:"availability zones JSON configuration for the bosh director"`
		NetworksConfiguration          string `short:"n"  long:"networks-configuration"  description:"complete network configuration for the bosh director"`
	}
}

//go:generate counterfeiter -o ./fakes/bosh_form_service.go --fake-name BoshFormService . boshFormService
type boshFormService interface {
	GetForm(path string) (api.Form, error)
	PostForm(api.PostFormInput) error
	AvailabilityZones() (map[string]string, error)
}

func NewConfigureBosh(s boshFormService, l logger) ConfigureBosh {
	return ConfigureBosh{service: s, logger: l}
}

func (c ConfigureBosh) Execute(args []string) error {
	_, err := flags.Parse(&c.Options, args)
	if err != nil {
		return err
	}

	if c.Options.IaaSConfiguration != "" {
		c.logger.Printf("configuring iaas specific options for bosh tile")
		err = c.configureForm(iaasConfigurationPath, c.Options.IaaSConfiguration)
		if err != nil {
			return err
		}
	}

	if c.Options.DirectorConfiguration != "" {
		c.logger.Printf("configuring director options for bosh tile")
		err = c.configureForm(directorConfigurationPath, c.Options.DirectorConfiguration)
		if err != nil {
			return err
		}
	}

	if c.Options.AvailabilityZonesConfiguration != "" {
		c.logger.Printf("configuring availability zones for bosh tile")
		err = c.configureForm(availabilityZonesConfigurationPath, c.Options.AvailabilityZonesConfiguration)
		if err != nil {
			return err
		}
	}

	if c.Options.NetworksConfiguration != "" {
		c.logger.Printf("configuring network options for bosh tile")
		if err != nil {
			panic(err)
		}

		err = c.configureNetworkForm(networksConfigurationPath, c.Options.NetworksConfiguration)
		if err != nil {
			return err
		}
	}

	if c.Options.SecurityConfiguration != "" {
		c.logger.Printf("configuring security options for bosh tile")
		err = c.configureForm(securityConfigurationPath, c.Options.SecurityConfiguration)
		if err != nil {
			return err
		}
	}

	c.logger.Printf("finished configuring bosh tile")
	return nil
}

func (c ConfigureBosh) configureForm(path, configuration string) error {
	form, err := c.service.GetForm(path)
	if err != nil {
		return fmt.Errorf("could not fetch form: %s", err)
	}

	var initialConfig BoshConfiguration
	err = json.NewDecoder(strings.NewReader(configuration)).Decode(&initialConfig)
	if err != nil {
		return fmt.Errorf("could not decode json: %s", err)
	}

	initialConfig.AuthenticityToken = form.AuthenticityToken
	initialConfig.Method = form.RailsMethod

	formValues, err := query.Values(initialConfig)
	if err != nil {
		return err // cannot be tested
	}

	err = c.service.PostForm(api.PostFormInput{Form: form, EncodedPayload: formValues.Encode()})
	if err != nil {
		return fmt.Errorf("tile failed to configure: %s", err)
	}

	return nil
}

func (c ConfigureBosh) configureNetworkForm(path, configuration string) error {
	form, err := c.service.GetForm(path)
	if err != nil {
		return fmt.Errorf("could not fetch form: %s", err)
	}

	var initialConfig NetworksConfiguration
	err = json.NewDecoder(strings.NewReader(configuration)).Decode(&initialConfig)
	if err != nil {
		return fmt.Errorf("could not decode json: %s", err)
	}

	azMap, err := c.service.AvailabilityZones()
	if err != nil {
		return fmt.Errorf("could not fetch availability zones: %s", err)
	}

	for n, network := range initialConfig {
		for s, subnet := range network.Subnets {
			for _, azName := range subnet.AvailabilityZones {
				if azGuid, ok := azMap[azName]; ok {
					initialConfig[n].Subnets[s].AvailabilityZoneGUIDs = append(initialConfig[n].Subnets[s].AvailabilityZoneGUIDs, azGuid)
				}
			}
		}
	}

	finalConfig := BoshNetworkForm{
		Networks: initialConfig,
		CommonConfiguration: CommonConfiguration{
			AuthenticityToken: form.AuthenticityToken,
			Method:            form.RailsMethod,
		},
	}

	formValues, err := query.Values(finalConfig)
	if err != nil {
		return err // cannot be tested
	}

	err = c.service.PostForm(api.PostFormInput{Form: form, EncodedPayload: formValues.Encode()})
	if err != nil {
		return fmt.Errorf("tile failed to configure: %s", err)
	}

	return nil
}

func (c ConfigureBosh) Usage() Usage {
	return Usage{
		Description:      "configures the bosh director that is deployed by the Ops Manager",
		ShortDescription: "configures Ops Manager deployed bosh director",
		Flags:            c.Options,
	}
}
