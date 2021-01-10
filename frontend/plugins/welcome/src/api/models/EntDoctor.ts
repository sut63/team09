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
    EntDoctorEdges,
    EntDoctorEdgesFromJSON,
    EntDoctorEdgesFromJSONTyped,
    EntDoctorEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDoctor
 */
export interface EntDoctor {
    /**
     * Address holds the value of the "address" field.
     * @type {string}
     * @memberof EntDoctor
     */
    address?: string;
    /**
     * Age holds the value of the "age" field.
     * @type {number}
     * @memberof EntDoctor
     */
    age?: number;
    /**
     * 
     * @type {EntDoctorEdges}
     * @memberof EntDoctor
     */
    edges?: EntDoctorEdges;
    /**
     * Educational holds the value of the "educational" field.
     * @type {string}
     * @memberof EntDoctor
     */
    educational?: string;
    /**
     * Email holds the value of the "email" field.
     * @type {string}
     * @memberof EntDoctor
     */
    email?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDoctor
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntDoctor
     */
    name?: string;
    /**
     * Pnumber holds the value of the "pnumber" field.
     * @type {number}
     * @memberof EntDoctor
     */
    pnumber?: number;
}

export function EntDoctorFromJSON(json: any): EntDoctor {
    return EntDoctorFromJSONTyped(json, false);
}

export function EntDoctorFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDoctor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'address': !exists(json, 'address') ? undefined : json['address'],
        'age': !exists(json, 'age') ? undefined : json['age'],
        'edges': !exists(json, 'edges') ? undefined : EntDoctorEdgesFromJSON(json['edges']),
        'educational': !exists(json, 'educational') ? undefined : json['educational'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'pnumber': !exists(json, 'pnumber') ? undefined : json['pnumber'],
    };
}

export function EntDoctorToJSON(value?: EntDoctor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'address': value.address,
        'age': value.age,
        'edges': EntDoctorEdgesToJSON(value.edges),
        'educational': value.educational,
        'email': value.email,
        'id': value.id,
        'name': value.name,
        'pnumber': value.pnumber,
    };
}


