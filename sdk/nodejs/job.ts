// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'dbtcloud:index/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Version number of DBT to use in this job
     */
    public readonly dbtVersion!: pulumi.Output<string | undefined>;
    /**
     * Job identifier that this job defers to
     */
    public readonly deferringJobId!: pulumi.Output<number | undefined>;
    /**
     * Environment ID to create the job in
     */
    public readonly environmentId!: pulumi.Output<number>;
    /**
     * List of commands to execute for the job
     */
    public readonly executeSteps!: pulumi.Output<string[]>;
    /**
     * Flag for whether the job should generate documentation
     */
    public readonly generateDocs!: pulumi.Output<boolean | undefined>;
    /**
     * Flag for whether the job is marked active or deleted
     */
    public readonly isActive!: pulumi.Output<boolean | undefined>;
    /**
     * Job name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of threads to use in the job
     */
    public readonly numThreads!: pulumi.Output<number | undefined>;
    /**
     * Project ID to create the job in
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * Flag for whether the job should run generate sources
     */
    public readonly runGenerateSources!: pulumi.Output<boolean | undefined>;
    /**
     * Custom cron expression for schedule
     */
    public readonly scheduleCron!: pulumi.Output<string | undefined>;
    /**
     * List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
     */
    public readonly scheduleDays!: pulumi.Output<number[] | undefined>;
    /**
     * List of hours to execute the job at if running on a schedule
     */
    public readonly scheduleHours!: pulumi.Output<number[] | undefined>;
    /**
     * Number of hours between job executions if running on a schedule
     */
    public readonly scheduleInterval!: pulumi.Output<number | undefined>;
    /**
     * Type of schedule to use, one of every_day/ days_of_week/ custom_cron
     */
    public readonly scheduleType!: pulumi.Output<string | undefined>;
    /**
     * Whether this job defers on a previous run of itself
     */
    public readonly selfDeferring!: pulumi.Output<boolean | undefined>;
    /**
     * Target name for the DBT profile
     */
    public readonly targetName!: pulumi.Output<string | undefined>;
    /**
     * Number of seconds to allow the job to run before timing out
     */
    public readonly timeoutSeconds!: pulumi.Output<number | undefined>;
    /**
     * Flags for which types of triggers to use, keys of github_webhook, git_provider_webhook, schedule, custom_branch_only
     */
    public readonly triggers!: pulumi.Output<{[key: string]: boolean}>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobState | undefined;
            resourceInputs["dbtVersion"] = state ? state.dbtVersion : undefined;
            resourceInputs["deferringJobId"] = state ? state.deferringJobId : undefined;
            resourceInputs["environmentId"] = state ? state.environmentId : undefined;
            resourceInputs["executeSteps"] = state ? state.executeSteps : undefined;
            resourceInputs["generateDocs"] = state ? state.generateDocs : undefined;
            resourceInputs["isActive"] = state ? state.isActive : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numThreads"] = state ? state.numThreads : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["runGenerateSources"] = state ? state.runGenerateSources : undefined;
            resourceInputs["scheduleCron"] = state ? state.scheduleCron : undefined;
            resourceInputs["scheduleDays"] = state ? state.scheduleDays : undefined;
            resourceInputs["scheduleHours"] = state ? state.scheduleHours : undefined;
            resourceInputs["scheduleInterval"] = state ? state.scheduleInterval : undefined;
            resourceInputs["scheduleType"] = state ? state.scheduleType : undefined;
            resourceInputs["selfDeferring"] = state ? state.selfDeferring : undefined;
            resourceInputs["targetName"] = state ? state.targetName : undefined;
            resourceInputs["timeoutSeconds"] = state ? state.timeoutSeconds : undefined;
            resourceInputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.executeSteps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executeSteps'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.triggers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggers'");
            }
            resourceInputs["dbtVersion"] = args ? args.dbtVersion : undefined;
            resourceInputs["deferringJobId"] = args ? args.deferringJobId : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["executeSteps"] = args ? args.executeSteps : undefined;
            resourceInputs["generateDocs"] = args ? args.generateDocs : undefined;
            resourceInputs["isActive"] = args ? args.isActive : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["numThreads"] = args ? args.numThreads : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["runGenerateSources"] = args ? args.runGenerateSources : undefined;
            resourceInputs["scheduleCron"] = args ? args.scheduleCron : undefined;
            resourceInputs["scheduleDays"] = args ? args.scheduleDays : undefined;
            resourceInputs["scheduleHours"] = args ? args.scheduleHours : undefined;
            resourceInputs["scheduleInterval"] = args ? args.scheduleInterval : undefined;
            resourceInputs["scheduleType"] = args ? args.scheduleType : undefined;
            resourceInputs["selfDeferring"] = args ? args.selfDeferring : undefined;
            resourceInputs["targetName"] = args ? args.targetName : undefined;
            resourceInputs["timeoutSeconds"] = args ? args.timeoutSeconds : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * Version number of DBT to use in this job
     */
    dbtVersion?: pulumi.Input<string>;
    /**
     * Job identifier that this job defers to
     */
    deferringJobId?: pulumi.Input<number>;
    /**
     * Environment ID to create the job in
     */
    environmentId?: pulumi.Input<number>;
    /**
     * List of commands to execute for the job
     */
    executeSteps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag for whether the job should generate documentation
     */
    generateDocs?: pulumi.Input<boolean>;
    /**
     * Flag for whether the job is marked active or deleted
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Job name
     */
    name?: pulumi.Input<string>;
    /**
     * Number of threads to use in the job
     */
    numThreads?: pulumi.Input<number>;
    /**
     * Project ID to create the job in
     */
    projectId?: pulumi.Input<number>;
    /**
     * Flag for whether the job should run generate sources
     */
    runGenerateSources?: pulumi.Input<boolean>;
    /**
     * Custom cron expression for schedule
     */
    scheduleCron?: pulumi.Input<string>;
    /**
     * List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
     */
    scheduleDays?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * List of hours to execute the job at if running on a schedule
     */
    scheduleHours?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Number of hours between job executions if running on a schedule
     */
    scheduleInterval?: pulumi.Input<number>;
    /**
     * Type of schedule to use, one of every_day/ days_of_week/ custom_cron
     */
    scheduleType?: pulumi.Input<string>;
    /**
     * Whether this job defers on a previous run of itself
     */
    selfDeferring?: pulumi.Input<boolean>;
    /**
     * Target name for the DBT profile
     */
    targetName?: pulumi.Input<string>;
    /**
     * Number of seconds to allow the job to run before timing out
     */
    timeoutSeconds?: pulumi.Input<number>;
    /**
     * Flags for which types of triggers to use, keys of github_webhook, git_provider_webhook, schedule, custom_branch_only
     */
    triggers?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Version number of DBT to use in this job
     */
    dbtVersion?: pulumi.Input<string>;
    /**
     * Job identifier that this job defers to
     */
    deferringJobId?: pulumi.Input<number>;
    /**
     * Environment ID to create the job in
     */
    environmentId: pulumi.Input<number>;
    /**
     * List of commands to execute for the job
     */
    executeSteps: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag for whether the job should generate documentation
     */
    generateDocs?: pulumi.Input<boolean>;
    /**
     * Flag for whether the job is marked active or deleted
     */
    isActive?: pulumi.Input<boolean>;
    /**
     * Job name
     */
    name?: pulumi.Input<string>;
    /**
     * Number of threads to use in the job
     */
    numThreads?: pulumi.Input<number>;
    /**
     * Project ID to create the job in
     */
    projectId: pulumi.Input<number>;
    /**
     * Flag for whether the job should run generate sources
     */
    runGenerateSources?: pulumi.Input<boolean>;
    /**
     * Custom cron expression for schedule
     */
    scheduleCron?: pulumi.Input<string>;
    /**
     * List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
     */
    scheduleDays?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * List of hours to execute the job at if running on a schedule
     */
    scheduleHours?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Number of hours between job executions if running on a schedule
     */
    scheduleInterval?: pulumi.Input<number>;
    /**
     * Type of schedule to use, one of every_day/ days_of_week/ custom_cron
     */
    scheduleType?: pulumi.Input<string>;
    /**
     * Whether this job defers on a previous run of itself
     */
    selfDeferring?: pulumi.Input<boolean>;
    /**
     * Target name for the DBT profile
     */
    targetName?: pulumi.Input<string>;
    /**
     * Number of seconds to allow the job to run before timing out
     */
    timeoutSeconds?: pulumi.Input<number>;
    /**
     * Flags for which types of triggers to use, keys of github_webhook, git_provider_webhook, schedule, custom_branch_only
     */
    triggers: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
}