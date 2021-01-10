/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentFromJSONTyped,
    EntDepartmentToJSON,
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorFromJSONTyped,
    EntDoctorToJSON,
    EntOffice,
    EntOfficeFromJSON,
    EntOfficeFromJSONTyped,
    EntOfficeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntScheduleEdges
 */
export interface EntScheduleEdges {
    /**
     * 
     * @type {EntDepartment}
     * @memberof EntScheduleEdges
     */
    department?: EntDepartment;
    /**
     * 
     * @type {EntDoctor}
     * @memberof EntScheduleEdges
     */
    docter?: EntDoctor;
    /**
     * 
     * @type {EntOffice}
     * @memberof EntScheduleEdges
     */
    office?: EntOffice;
}

export function EntScheduleEdgesFromJSON(json: any): EntScheduleEdges {
    return EntScheduleEdgesFromJSONTyped(json, false);
}

export function EntScheduleEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntScheduleEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'department': !exists(json, 'department') ? undefined : EntDepartmentFromJSON(json['department']),
        'docter': !exists(json, 'docter') ? undefined : EntDoctorFromJSON(json['docter']),
        'office': !exists(json, 'office') ? undefined : EntOfficeFromJSON(json['office']),
    };
}

export function EntScheduleEdgesToJSON(value?: EntScheduleEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'department': EntDepartmentToJSON(value.department),
        'docter': EntDoctorToJSON(value.docter),
        'office': EntOfficeToJSON(value.office),
    };
}

