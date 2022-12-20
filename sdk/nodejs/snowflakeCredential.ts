// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SnowflakeCredential extends pulumi.CustomResource {
    /**
     * Get an existing SnowflakeCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnowflakeCredentialState, opts?: pulumi.CustomResourceOptions): SnowflakeCredential {
        return new SnowflakeCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/snowflakeCredential:SnowflakeCredential';

    /**
     * Returns true if the given object is an instance of SnowflakeCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnowflakeCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnowflakeCredential.__pulumiType;
    }

    /**
     * The type of Snowflake credential ('password' or 'keypair')
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * The system Snowflake credential ID
     */
    public /*out*/ readonly credentialId!: pulumi.Output<number>;
    /**
     * Database to connect to
     */
    public readonly database!: pulumi.Output<string | undefined>;
    /**
     * Whether the Snowflake credential is active
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * Number of threads to use
     */
    public readonly numThreads!: pulumi.Output<number>;
    /**
     * Password for Snowflake
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Private key for Snowflake
     */
    public readonly privateKey!: pulumi.Output<string | undefined>;
    /**
     * Private key passphrase for Snowflake
     */
    public readonly privateKeyPassphrase!: pulumi.Output<string | undefined>;
    /**
     * Project ID to create the Snowflake credential in
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Role to assume
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * Default schema name
     */
    public readonly schema!: pulumi.Output<string>;
    /**
     * Username for Snowflake
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * Warehouse to use
     */
    public readonly warehouse!: pulumi.Output<string | undefined>;

    /**
     * Create a SnowflakeCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnowflakeCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnowflakeCredentialArgs | SnowflakeCredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnowflakeCredentialState | undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["credentialId"] = state ? state.credentialId : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["numThreads"] = state ? state.numThreads : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["privateKeyPassphrase"] = state ? state.privateKeyPassphrase : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["warehouse"] = state ? state.warehouse : undefined;
        } else {
            const args = argsOrState as SnowflakeCredentialArgs | undefined;
            if ((!args || args.authType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authType'");
            }
            if ((!args || args.numThreads === undefined) && !opts.urn) {
                throw new Error("Missing required property 'numThreads'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.schema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schema'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["numThreads"] = args ? args.numThreads : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["privateKeyPassphrase"] = args ? args.privateKeyPassphrase : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["warehouse"] = args ? args.warehouse : undefined;
            resourceInputs["credentialId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnowflakeCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnowflakeCredential resources.
 */
export interface SnowflakeCredentialState {
    /**
     * The type of Snowflake credential ('password' or 'keypair')
     */
    authType?: pulumi.Input<string>;
    /**
     * The system Snowflake credential ID
     */
    credentialId?: pulumi.Input<number>;
    /**
     * Database to connect to
     */
    database?: pulumi.Input<string>;
    /**
     * Whether the Snowflake credential is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Number of threads to use
     */
    numThreads?: pulumi.Input<number>;
    /**
     * Password for Snowflake
     */
    password?: pulumi.Input<string>;
    /**
     * Private key for Snowflake
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key passphrase for Snowflake
     */
    privateKeyPassphrase?: pulumi.Input<string>;
    /**
     * Project ID to create the Snowflake credential in
     */
    projectId?: pulumi.Input<number>;
    /**
     * Role to assume
     */
    role?: pulumi.Input<string>;
    /**
     * Default schema name
     */
    schema?: pulumi.Input<string>;
    /**
     * Username for Snowflake
     */
    user?: pulumi.Input<string>;
    /**
     * Warehouse to use
     */
    warehouse?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnowflakeCredential resource.
 */
export interface SnowflakeCredentialArgs {
    /**
     * The type of Snowflake credential ('password' or 'keypair')
     */
    authType: pulumi.Input<string>;
    /**
     * Database to connect to
     */
    database?: pulumi.Input<string>;
    /**
     * Whether the Snowflake credential is active
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Number of threads to use
     */
    numThreads: pulumi.Input<number>;
    /**
     * Password for Snowflake
     */
    password?: pulumi.Input<string>;
    /**
     * Private key for Snowflake
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Private key passphrase for Snowflake
     */
    privateKeyPassphrase?: pulumi.Input<string>;
    /**
     * Project ID to create the Snowflake credential in
     */
    projectId: pulumi.Input<number>;
    /**
     * Role to assume
     */
    role?: pulumi.Input<string>;
    /**
     * Default schema name
     */
    schema: pulumi.Input<string>;
    /**
     * Username for Snowflake
     */
    user: pulumi.Input<string>;
    /**
     * Warehouse to use
     */
    warehouse?: pulumi.Input<string>;
}