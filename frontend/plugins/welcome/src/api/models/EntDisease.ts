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
    EntDiseaseEdges,
    EntDiseaseEdgesFromJSON,
    EntDiseaseEdgesFromJSONTyped,
    EntDiseaseEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDisease
 */
export interface EntDisease {
    /**
     * Disease holds the value of the "disease" field.
     * @type {string}
     * @memberof EntDisease
     */
    disease?: string;
    /**
     * 
     * @type {EntDiseaseEdges}
     * @memberof EntDisease
     */
    edges?: EntDiseaseEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDisease
     */
    id?: number;
}

export function EntDiseaseFromJSON(json: any): EntDisease {
    return EntDiseaseFromJSONTyped(json, false);
}

export function EntDiseaseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDisease {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'disease': !exists(json, 'disease') ? undefined : json['disease'],
        'edges': !exists(json, 'edges') ? undefined : EntDiseaseEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDiseaseToJSON(value?: EntDisease | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'disease': value.disease,
        'edges': EntDiseaseEdgesToJSON(value.edges),
        'id': value.id,
    };
}

