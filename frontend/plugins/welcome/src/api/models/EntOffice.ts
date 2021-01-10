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
    EntOfficeEdges,
    EntOfficeEdgesFromJSON,
    EntOfficeEdgesFromJSONTyped,
    EntOfficeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOffice
 */
export interface EntOffice {
    /**
     * AddedTime1 holds the value of the "added_time1" field.
     * @type {string}
     * @memberof EntOffice
     */
    addedTime1?: string;
    /**
     * AddedTime2 holds the value of the "added_time2" field.
     * @type {string}
     * @memberof EntOffice
     */
    addedTime2?: string;
    /**
     * 
     * @type {EntOfficeEdges}
     * @memberof EntOffice
     */
    edges?: EntOfficeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntOffice
     */
    id?: number;
    /**
     * Officename holds the value of the "officename" field.
     * @type {string}
     * @memberof EntOffice
     */
    officename?: string;
}

export function EntOfficeFromJSON(json: any): EntOffice {
    return EntOfficeFromJSONTyped(json, false);
}

export function EntOfficeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOffice {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime1': !exists(json, 'added_time1') ? undefined : json['added_time1'],
        'addedTime2': !exists(json, 'added_time2') ? undefined : json['added_time2'],
        'edges': !exists(json, 'edges') ? undefined : EntOfficeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'officename': !exists(json, 'officename') ? undefined : json['officename'],
    };
}

export function EntOfficeToJSON(value?: EntOffice | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added_time1': value.addedTime1,
        'added_time2': value.addedTime2,
        'edges': EntOfficeEdgesToJSON(value.edges),
        'id': value.id,
        'officename': value.officename,
    };
}


