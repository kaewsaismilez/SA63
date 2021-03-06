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
    EntMedicalfile,
    EntMedicalfileFromJSON,
    EntMedicalfileFromJSONTyped,
    EntMedicalfileToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEmployeeEdges
 */
export interface EntEmployeeEdges {
    /**
     * Medicalfiles holds the value of the medicalfiles edge.
     * @type {Array<EntMedicalfile>}
     * @memberof EntEmployeeEdges
     */
    medicalfiles?: Array<EntMedicalfile>;
}

export function EntEmployeeEdgesFromJSON(json: any): EntEmployeeEdges {
    return EntEmployeeEdgesFromJSONTyped(json, false);
}

export function EntEmployeeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEmployeeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'medicalfiles': !exists(json, 'medicalfiles') ? undefined : ((json['medicalfiles'] as Array<any>).map(EntMedicalfileFromJSON)),
    };
}

export function EntEmployeeEdgesToJSON(value?: EntEmployeeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'medicalfiles': value.medicalfiles === undefined ? undefined : ((value.medicalfiles as Array<any>).map(EntMedicalfileToJSON)),
    };
}


