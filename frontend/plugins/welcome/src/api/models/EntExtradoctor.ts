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
    EntExtradoctorEdges,
    EntExtradoctorEdgesFromJSON,
    EntExtradoctorEdgesFromJSONTyped,
    EntExtradoctorEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntExtradoctor
 */
export interface EntExtradoctor {
    /**
     * 
     * @type {EntExtradoctorEdges}
     * @memberof EntExtradoctor
     */
    edges?: EntExtradoctorEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntExtradoctor
     */
    id?: number;
    /**
     * Specialname holds the value of the "specialname" field.
     * @type {string}
     * @memberof EntExtradoctor
     */
    specialname?: string;
}

export function EntExtradoctorFromJSON(json: any): EntExtradoctor {
    return EntExtradoctorFromJSONTyped(json, false);
}

export function EntExtradoctorFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntExtradoctor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntExtradoctorEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'specialname': !exists(json, 'specialname') ? undefined : json['specialname'],
    };
}

export function EntExtradoctorToJSON(value?: EntExtradoctor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntExtradoctorEdgesToJSON(value.edges),
        'id': value.id,
        'specialname': value.specialname,
    };
}


