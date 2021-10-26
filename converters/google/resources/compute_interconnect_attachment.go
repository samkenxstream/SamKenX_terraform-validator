// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
)

const ComputeInterconnectAttachmentAssetType string = "compute.googleapis.com/InterconnectAttachment"

func resourceConverterComputeInterconnectAttachment() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeInterconnectAttachmentAssetType,
		Convert:   GetComputeInterconnectAttachmentCaiObject,
	}
}

func GetComputeInterconnectAttachmentCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeInterconnectAttachmentApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeInterconnectAttachmentAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "InterconnectAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeInterconnectAttachmentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	adminEnabledProp, err := expandComputeInterconnectAttachmentAdminEnabled(d.Get("admin_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("admin_enabled"); ok || !reflect.DeepEqual(v, adminEnabledProp) {
		obj["adminEnabled"] = adminEnabledProp
	}
	interconnectProp, err := expandComputeInterconnectAttachmentInterconnect(d.Get("interconnect"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("interconnect"); !isEmptyValue(reflect.ValueOf(interconnectProp)) && (ok || !reflect.DeepEqual(v, interconnectProp)) {
		obj["interconnect"] = interconnectProp
	}
	descriptionProp, err := expandComputeInterconnectAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	mtuProp, err := expandComputeInterconnectAttachmentMtu(d.Get("mtu"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mtu"); !isEmptyValue(reflect.ValueOf(mtuProp)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}
	bandwidthProp, err := expandComputeInterconnectAttachmentBandwidth(d.Get("bandwidth"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bandwidth"); !isEmptyValue(reflect.ValueOf(bandwidthProp)) && (ok || !reflect.DeepEqual(v, bandwidthProp)) {
		obj["bandwidth"] = bandwidthProp
	}
	edgeAvailabilityDomainProp, err := expandComputeInterconnectAttachmentEdgeAvailabilityDomain(d.Get("edge_availability_domain"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("edge_availability_domain"); !isEmptyValue(reflect.ValueOf(edgeAvailabilityDomainProp)) && (ok || !reflect.DeepEqual(v, edgeAvailabilityDomainProp)) {
		obj["edgeAvailabilityDomain"] = edgeAvailabilityDomainProp
	}
	typeProp, err := expandComputeInterconnectAttachmentType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	routerProp, err := expandComputeInterconnectAttachmentRouter(d.Get("router"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	nameProp, err := expandComputeInterconnectAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	candidateSubnetsProp, err := expandComputeInterconnectAttachmentCandidateSubnets(d.Get("candidate_subnets"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("candidate_subnets"); !isEmptyValue(reflect.ValueOf(candidateSubnetsProp)) && (ok || !reflect.DeepEqual(v, candidateSubnetsProp)) {
		obj["candidateSubnets"] = candidateSubnetsProp
	}
	vlanTag8021qProp, err := expandComputeInterconnectAttachmentVlanTag8021q(d.Get("vlan_tag8021q"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vlan_tag8021q"); !isEmptyValue(reflect.ValueOf(vlanTag8021qProp)) && (ok || !reflect.DeepEqual(v, vlanTag8021qProp)) {
		obj["vlanTag8021q"] = vlanTag8021qProp
	}
	ipsecInternalAddressesProp, err := expandComputeInterconnectAttachmentIpsecInternalAddresses(d.Get("ipsec_internal_addresses"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ipsec_internal_addresses"); !isEmptyValue(reflect.ValueOf(ipsecInternalAddressesProp)) && (ok || !reflect.DeepEqual(v, ipsecInternalAddressesProp)) {
		obj["ipsecInternalAddresses"] = ipsecInternalAddressesProp
	}
	encryptionProp, err := expandComputeInterconnectAttachmentEncryption(d.Get("encryption"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption"); !isEmptyValue(reflect.ValueOf(encryptionProp)) && (ok || !reflect.DeepEqual(v, encryptionProp)) {
		obj["encryption"] = encryptionProp
	}
	regionProp, err := expandComputeInterconnectAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeInterconnectAttachmentAdminEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentInterconnect(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentMtu(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentBandwidth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentEdgeAvailabilityDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeInterconnectAttachmentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentCandidateSubnets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentVlanTag8021q(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentIpsecInternalAddresses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for ipsec_internal_addresses: nil")
		}
		f, err := parseRegionalFieldValue("addresses", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ipsec_internal_addresses: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeInterconnectAttachmentEncryption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}