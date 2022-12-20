// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Dbtcloud
{
    public static class GetConnection
    {
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("dbtcloud:index/getConnection:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("dbtcloud:index/getConnection:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public int ConnectionId { get; set; }

        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public GetConnectionArgs()
        {
        }
        public static new GetConnectionArgs Empty => new GetConnectionArgs();
    }

    public sealed class GetConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public Input<int> ConnectionId { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public GetConnectionInvokeArgs()
        {
        }
        public static new GetConnectionInvokeArgs Empty => new GetConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        public readonly string Account;
        public readonly bool AllowKeepAlive;
        public readonly bool AllowSso;
        public readonly int ConnectionId;
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsActive;
        public readonly string Name;
        public readonly int ProjectId;
        public readonly string Role;
        public readonly string Type;
        public readonly string Warehouse;

        [OutputConstructor]
        private GetConnectionResult(
            string account,

            bool allowKeepAlive,

            bool allowSso,

            int connectionId,

            string database,

            string id,

            bool isActive,

            string name,

            int projectId,

            string role,

            string type,

            string warehouse)
        {
            Account = account;
            AllowKeepAlive = allowKeepAlive;
            AllowSso = allowSso;
            ConnectionId = connectionId;
            Database = database;
            Id = id;
            IsActive = isActive;
            Name = name;
            ProjectId = projectId;
            Role = role;
            Type = type;
            Warehouse = warehouse;
        }
    }
}
