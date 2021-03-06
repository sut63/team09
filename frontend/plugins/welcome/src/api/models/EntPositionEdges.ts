/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
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
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorFromJSONTyped,
    EntDoctorToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPositionEdges
 */
export interface EntPositionEdges {
    /**
     * Doctors holds the value of the doctors edge.
     * @type {Array<EntDoctor>}
     * @memberof EntPositionEdges
     */
    doctors?: Array<EntDoctor>;
}

export function EntPositionEdgesFromJSON(json: any): EntPositionEdges {
    return EntPositionEdgesFromJSONTyped(json, false);
}

export function EntPositionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPositionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctors': !exists(json, 'doctors') ? undefined : ((json['doctors'] as Array<any>).map(EntDoctorFromJSON)),
    };
}

export function EntPositionEdgesToJSON(value?: EntPositionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctors': value.doctors === undefined ? undefined : ((value.doctors as Array<any>).map(EntDoctorToJSON)),
    };
}


