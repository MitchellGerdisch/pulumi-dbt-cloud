// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getConnection(args: GetConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("dbtcloud:index/getConnection:getConnection", {
        "connectionId": args.connectionId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getConnection.
 */
export interface GetConnectionArgs {
    connectionId: number;
    projectId: number;
}

/**
 * A collection of values returned by getConnection.
 */
export interface GetConnectionResult {
    readonly account: string;
    readonly allowKeepAlive: boolean;
    readonly allowSso: boolean;
    readonly connectionId: number;
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isActive: boolean;
    readonly name: string;
    readonly projectId: number;
    readonly role: string;
    readonly type: string;
    readonly warehouse: string;
}

export function getConnectionOutput(args: GetConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionResult> {
    return pulumi.output(args).apply(a => getConnection(a, opts))
}

/**
 * A collection of arguments for invoking getConnection.
 */
export interface GetConnectionOutputArgs {
    connectionId: pulumi.Input<number>;
    projectId: pulumi.Input<number>;
}
