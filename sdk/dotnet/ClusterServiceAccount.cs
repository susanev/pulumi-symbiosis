// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Kuraudo.Symbiosis
{
    /// <summary>
    /// Manages Kubernetes clusters service accounts.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Symbiosis = Kuraudo.Symbiosis;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Symbiosis.ClusterServiceAccount("example", new()
    ///     {
    ///         ClusterName = symbiosis_cluster.Example.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SymbiosisResourceType("symbiosis:index/clusterServiceAccount:ClusterServiceAccount")]
    public partial class ClusterServiceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster CA certificate
        /// </summary>
        [Output("clusterCaCertificate")]
        public Output<string> ClusterCaCertificate { get; private set; } = null!;

        /// <summary>
        /// Cluster name. Changing the name forces re-creation.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Service account token
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterServiceAccount(string name, ClusterServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("symbiosis:index/clusterServiceAccount:ClusterServiceAccount", name, args ?? new ClusterServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterServiceAccount(string name, Input<string> id, ClusterServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("symbiosis:index/clusterServiceAccount:ClusterServiceAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/kuraudo-io/pulumi-symbiosis/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterServiceAccount Get(string name, Input<string> id, ClusterServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterServiceAccount(name, id, state, options);
        }
    }

    public sealed class ClusterServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster name. Changing the name forces re-creation.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        public ClusterServiceAccountArgs()
        {
        }
        public static new ClusterServiceAccountArgs Empty => new ClusterServiceAccountArgs();
    }

    public sealed class ClusterServiceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster CA certificate
        /// </summary>
        [Input("clusterCaCertificate")]
        public Input<string>? ClusterCaCertificate { get; set; }

        /// <summary>
        /// Cluster name. Changing the name forces re-creation.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Service account token
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public ClusterServiceAccountState()
        {
        }
        public static new ClusterServiceAccountState Empty => new ClusterServiceAccountState();
    }
}