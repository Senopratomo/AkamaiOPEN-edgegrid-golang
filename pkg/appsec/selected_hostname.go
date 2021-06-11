package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The SelectedHostname interface supports retrieving and modifying the list of hostnames protected under
	// a configuration.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#selectedhostnames
	SelectedHostname interface {
		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getselectedhostnames
		GetSelectedHostnames(ctx context.Context, params GetSelectedHostnamesRequest) (*GetSelectedHostnamesResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getselectedhostnames
		GetSelectedHostname(ctx context.Context, params GetSelectedHostnameRequest) (*GetSelectedHostnameResponse, error)

		// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putselectedhostnames
		UpdateSelectedHostname(ctx context.Context, params UpdateSelectedHostnameRequest) (*UpdateSelectedHostnameResponse, error)
	}

	// GetSelectedHostnamesRequest is used to retrieve the selected hostnames for a configuration.
	GetSelectedHostnamesRequest struct {
		ConfigID     int        `json:"configId"`
		Version      int        `json:"version"`
		HostnameList []Hostname `json:"hostnameList"`
	}

	// GetSelectedHostnamesResponse is returned from a call to GetSelectedHostnames.
	GetSelectedHostnamesResponse struct {
		HostnameList []Hostname `json:"hostnameList,omitempty"`
	}

	// GetSelectedHostnameRequest is used to retrieve the selected hostnames for a configuration.
	GetSelectedHostnameRequest struct {
		ConfigID     int        `json:"configId"`
		Version      int        `json:"version"`
		HostnameList []Hostname `json:"hostnameList"`
	}

	// GetSelectedHostnameResponse is returned from a call to GetSelectedHostname.
	GetSelectedHostnameResponse struct {
		HostnameList []Hostname `json:"hostnameList"`
	}

	// UpdateSelectedHostnameRequest is used to modify the selected hostnames for a configuration.
	UpdateSelectedHostnameRequest struct {
		ConfigID     int        `json:"configId"`
		Version      int        `json:"version"`
		HostnameList []Hostname `json:"hostnameList"`
	}

	// UpdateSelectedHostnameResponse is returned from a call to UpdateSelectedHostname.
	UpdateSelectedHostnameResponse struct {
		HostnameList []Hostname `json:"hostnameList"`
	}

	// Hostname describes a hostname that may be protected.
	Hostname struct {
		Hostname string `json:"hostname"`
	}
)

// Validate validates a GetSelectedHostnameRequest.
func (v GetSelectedHostnameRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates a GetSelectedHostnamesRequest.
func (v GetSelectedHostnamesRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

// Validate validates an UpdateSelectedHostnameRequest.
func (v UpdateSelectedHostnameRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
	}.Filter()
}

func (p *appsec) GetSelectedHostname(ctx context.Context, params GetSelectedHostnameRequest) (*GetSelectedHostnameResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetSelectedHostname")

	var rval GetSelectedHostnameResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/selected-hostnames",
		params.ConfigID,
		params.Version)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetSelectedHostname request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetSelectedHostname request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) GetSelectedHostnames(ctx context.Context, params GetSelectedHostnamesRequest) (*GetSelectedHostnamesResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetSelectedHostnames")

	var rval GetSelectedHostnamesResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/selected-hostnames",
		params.ConfigID,
		params.Version)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetSelectedHostnames request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("GetSelectedHostnames request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) UpdateSelectedHostname(ctx context.Context, params UpdateSelectedHostnameRequest) (*UpdateSelectedHostnameResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateSelectedHostname")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/selected-hostnames",
		params.ConfigID,
		params.Version,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateSelectedHostname request: %w", err)
	}

	var rval UpdateSelectedHostnameResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("UpdateSelectedHostname request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}
