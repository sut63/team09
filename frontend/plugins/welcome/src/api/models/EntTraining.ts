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
    EntTrainingEdges,
    EntTrainingEdgesFromJSON,
    EntTrainingEdgesFromJSONTyped,
    EntTrainingEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTraining
 */
export interface EntTraining {
    /**
     * Branch holds the value of the "branch" field.
     * @type {string}
     * @memberof EntTraining
     */
    branch?: string;
    /**
     * Dateone holds the value of the "dateone" field.
     * @type {string}
     * @memberof EntTraining
     */
    dateone?: string;
    /**
     * Datetwo holds the value of the "datetwo" field.
     * @type {string}
     * @memberof EntTraining
     */
    datetwo?: string;
    /**
     * 
     * @type {EntTrainingEdges}
     * @memberof EntTraining
     */
    edges?: EntTrainingEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTraining
     */
    id?: number;
}

export function EntTrainingFromJSON(json: any): EntTraining {
    return EntTrainingFromJSONTyped(json, false);
}

export function EntTrainingFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTraining {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'branch': !exists(json, 'branch') ? undefined : json['branch'],
        'dateone': !exists(json, 'dateone') ? undefined : json['dateone'],
        'datetwo': !exists(json, 'datetwo') ? undefined : json['datetwo'],
        'edges': !exists(json, 'edges') ? undefined : EntTrainingEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntTrainingToJSON(value?: EntTraining | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'branch': value.branch,
        'dateone': value.dateone,
        'datetwo': value.datetwo,
        'edges': EntTrainingEdgesToJSON(value.edges),
        'id': value.id,
    };
}

