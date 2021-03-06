// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

// JobScheduleTypeEnum Enum with underlying type: string
type JobScheduleTypeEnum string

// Set of constants representing the allowable values for JobScheduleTypeEnum
const (
	JobScheduleTypeScheduled JobScheduleTypeEnum = "SCHEDULED"
	JobScheduleTypeImmediate JobScheduleTypeEnum = "IMMEDIATE"
)

var mappingJobScheduleType = map[string]JobScheduleTypeEnum{
	"SCHEDULED": JobScheduleTypeScheduled,
	"IMMEDIATE": JobScheduleTypeImmediate,
}

// GetJobScheduleTypeEnumValues Enumerates the set of values for JobScheduleTypeEnum
func GetJobScheduleTypeEnumValues() []JobScheduleTypeEnum {
	values := make([]JobScheduleTypeEnum, 0)
	for _, v := range mappingJobScheduleType {
		values = append(values, v)
	}
	return values
}
