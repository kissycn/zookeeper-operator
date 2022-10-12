# zookeeper-operator
Apache ZooKeeper provides a reliable, centralized register of configuration data and services for distributed applications.

[Overview of Apache ZooKeeper](https://zookeeper.apache.org)

The software provides full life cycle management of Zookeeper in the kubernetes environment.
Use tools CRD and kube-builder to implement the cluster installation and deployment of zookeeper in the kubernetes environment, monitor alarms, log collection, etc.

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/zookeeper-operator:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/zookeeper-operator:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Parameters

### Common parameters
| Name                           | Description                                                                                      | Value |
|--------------------------------| ------------------------------------------------------------------------------------------------ |------|
| `commonLabels`                 | Add labels to all the deployed resources                                                         | `{}` |
| `commonAnnotations`            | Add annotations to all the deployed resources                                                    | `{}` |


### ZooKeeper chart parameters

| Name                               | Description                                                                                                                                                        | Value                                              |
|------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| `conf.dataDir`                     | Dedicated data directory                                                                                                                                           | `"/opt/dtweave/zookeeper/data"`                    |
| `conf.dataLogDir`                  | Dedicated data log directory                                                                                                                                       | `"/opt/dtweave/zookeeper/data_log"`                |
| `conf.tickTime`                    | Basic time unit (in milliseconds) used by ZooKeeper for heartbeats                                                                                                 | `2000`                                             |
| `conf.initLimit`                   | ZooKeeper uses to limit the length of time the ZooKeeper servers in quorum have to connect to a leader                                                             | `10`                                               |
| `conf.syncLimit`                   | How far out of date a server can be from a leader                                                                                                                  | `5`                                                |
| `conf.maxClientCnxns`              | Limits the number of concurrent connections that a single client may make to a single member of the ZooKeeper ensemble                                             | `60`                                               |
| `conf.maxSessionTimeout`           | Maximum session timeout (in milliseconds) that the server will allow the client to negotiate                                                                       | `40000`                                            |
| `conf.autopurge.snapRetainCount`   | The most recent snapshots amount (and corresponding transaction logs) to retain                                                                                    | `3`                                                |
| `conf.autopurge.purgeInterval`     | The time interval (in hours) for which the purge task has to be triggered                                                                                          | `0`                                                |
| `conf.logLevel`                    | Log level for the ZooKeeper server. ERROR by default                                                                                                               | `ERROR`                                            |
| `conf.jvmFlags`                    | Default JVM flags for the ZooKeeper process                                                                                                                        | `""`                                               |
| `conf.existingCfgConfigmap`        | The name of an existing ConfigMap with your custom configuration for ZooKeeper zoo.cfg file<br/>Notice：If set this value operator will replace all other conf item | `""`                                               |
| `conf.whitelistCommands`           | zookeeper exec command                                                                                                                                             | `"cons, envi, conf, crst, srvr, stat, mntr, ruok"` |
| `conf.additionalConfig`            | key-value map of additional zookeeper configuration parameters                                                                                                     | {}                                                 |
| `conf.extraEnvVars`                | Array with extra environment variables to add to ZooKeeper nodes                                                                                                   | `[]`                                               |
| `conf.auth.client.enabled`         | Enable ZooKeeper client-server authentication. It uses SASL/Digest-MD5                                                                                             | `false`                                            |
| `conf.auth.client.clientUser`      | User that will use ZooKeeper clients to auth                                                                                                                       | `""`                                               |
| `conf.auth.client.clientPassword`  | Password that will use ZooKeeper clients to auth                                                                                                                   | `""`                                               |
| `conf.auth.client.serverUsers`     | Comma, semicolon or whitespace separated list of user to be created                                                                                                | `""`                                               |
| `conf.auth.client.serverPasswords` | Comma, semicolon or whitespace separated list of passwords to assign to users when created                                                                         | `""`                                               |
| `conf.auth.client.existingSecret`  | Use existing secret (ignores previous passwords)                                                                                                                   | `""`                                               |
| `conf.auth.quorum.enabled`         | Enable ZooKeeper server-server authentication. It uses SASL/Digest-MD5                                                                                             | `false`                                            |
| `conf.auth.quorum.learnerUser`     | User that the ZooKeeper quorumLearner will use to authenticate to quorumServers.                                                                                   | `""`                                               |
| `conf.auth.quorum.learnerPassword` | Password that the ZooKeeper quorumLearner will use to authenticate to quorumServers.                                                                               | `""`                                               |
| `conf.auth.quorum.serverUsers`     | Comma, semicolon or whitespace separated list of users for the quorumServers.                                                                                      | `""`                                               |
| `conf.auth.quorum.serverPasswords` | Comma, semicolon or whitespace separated list of passwords to assign to users when created                                                                         | `""`                                               |
| `conf.auth.quorum.existingSecret`  | Use existing secret (ignores previous passwords)                                                                                                                   | `""`                                               |

### Statefulset PodPolicy parameters

| Name                                 | Description                                                                                                                                                                                      | Value               |
|--------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------|
| `replicasCount`                      | Number of ZooKeeper nodes                                                                                                                                                                        | `1`                 |
| `podManagementPolicy`                | StatefulSet controller supports relax its ordering guarantees while preserving its uniqueness and identity guarantees. There are two valid pod management policies: `OrderedReady` and `Parallel` | `Parallel`          |
| `podLabels`                          | Extra labels for ZooKeeper pods                                                                                                                                                                  | `{}`                |
| `podAnnotations`                     | Annotations for ZooKeeper pods                                                                                                                                                                   | `{}`                |
| `hostAliases`                        | ZooKeeper pods host aliases                                                                                                                                                                      | `[]`                |
| `affinity`                           | Affinity for pod assignment                                                                                                                                                                      | `{}`                |
| `nodeSelector`                       | Node labels for pod assignment                                                                                                                                                                   | `{}`                |
| `tolerations`                        | Tolerations for pod assignment                                                                                                                                                                   | `[]`                |
| `topologySpreadConstraints`          | Topology Spread Constraints for pod assignment spread across your cluster among failure-domains. Evaluated as a template                                                                         | `[]`                |
| `priorityClassName`                  | Name of the existing priority class to be used by ZooKeeper pods, priority class needs to be created beforehand                                                                                  | `""`                |
| `schedulerName`                      | Kubernetes pod scheduler registry                                                                                                                                                                | `""`                |
| `image.registry`                     | ZooKeeper image registry                                                                                                                                                                         | `docker.io`         |
| `image.repository`                   | ZooKeeper image repository                                                                                                                                                                       | `dtweave/zookeeper` |
| `image.tag`                          | ZooKeeper image tag (immutable tags are recommended)                                                                                                                                             | `v1.0.0-3.7.1`      |
| `image.digest`                       | ZooKeeper image digest in the way sha256:aa.... Please note this parameter, if set, will override the tag                                                                                        | `""`                |
| `image.pullPolicy`                   | ZooKeeper image pull policy                                                                                                                                                                      | `IfNotPresent`      |
| `image.pullSecrets`                  | Specify docker-registry secret names as an array                                                                                                                                                 | `[]`                |
| `containerPorts.client`              | ZooKeeper client container port                                                                                                                                                                  | `2181`              |
| `containerPorts.tls`                 | ZooKeeper TLS container port                                                                                                                                                                     | `3181`              |
| `containerPorts.follower`            | ZooKeeper follower container port                                                                                                                                                                | `2888`              |
| `containerPorts.election`            | ZooKeeper election container port                                                                                                                                                                | `3888`              |
| `livenessProbe.enabled`              | Enable livenessProbe on ZooKeeper containers                                                                                                                                                     | `true`              |
| `livenessProbe.initialDelaySeconds`  | Initial delay seconds for livenessProbe                                                                                                                                                          | `30`                |
| `livenessProbe.periodSeconds`        | Period seconds for livenessProbe                                                                                                                                                                 | `10`                |
| `livenessProbe.timeoutSeconds`       | Timeout seconds for livenessProbe                                                                                                                                                                | `5`                 |
| `livenessProbe.failureThreshold`     | Failure threshold for livenessProbe                                                                                                                                                              | `6`                 |
| `livenessProbe.successThreshold`     | Success threshold for livenessProbe                                                                                                                                                              | `1`                 |
| `livenessProbe.probeCommandTimeout`  | Probe command timeout for livenessProbe                                                                                                                                                          | `2`                 |
| `readinessProbe.enabled`             | Enable readinessProbe on ZooKeeper containers                                                                                                                                                    | `true`              |
| `readinessProbe.initialDelaySeconds` | Initial delay seconds for readinessProbe                                                                                                                                                         | `5`                 |
| `readinessProbe.periodSeconds`       | Period seconds for readinessProbe                                                                                                                                                                | `10`                |
| `readinessProbe.timeoutSeconds`      | Timeout seconds for readinessProbe                                                                                                                                                               | `5`                 |
| `readinessProbe.failureThreshold`    | Failure threshold for readinessProbe                                                                                                                                                             | `6`                 |
| `readinessProbe.successThreshold`    | Success threshold for readinessProbe                                                                                                                                                             | `1`                 |
| `readinessProbe.probeCommandTimeout` | Probe command timeout for readinessProbe                                                                                                                                                         | `2`                 |
| `startupProbe.enabled`               | Enable startupProbe on ZooKeeper containers                                                                                                                                                      | `false`             |
| `startupProbe.initialDelaySeconds`   | Initial delay seconds for startupProbe                                                                                                                                                           | `30`                |
| `startupProbe.periodSeconds`         | Period seconds for startupProbe                                                                                                                                                                  | `10`                |
| `startupProbe.timeoutSeconds`        | Timeout seconds for startupProbe                                                                                                                                                                 | `1`                 |
| `startupProbe.failureThreshold`      | Failure threshold for startupProbe                                                                                                                                                               | `15`                |
| `startupProbe.successThreshold`      | Success threshold for startupProbe                                                                                                                                                               | `1`                 |
| `lifecycleHooks`                     | for the ZooKeeper container(s) to automate configuration before or after startup                                                                                                                 | `{}`                |
| `resources.limits.cpu`               | The resources cpu for the ZooKeeper containers                                                                                                                                                   | `512m`              |
| `resources.limits.memory`            | The resources memory for the ZooKeeper containers                                                                                                                                          | `1Gi`               |
| `resources.requests.cpu`             | The requested cpu for the ZooKeeper containers                                                                                                                                                   | `250m`              |
| `resources.requests.memory`          | The requested memory for the ZooKeeper containers                                                                                                                                                | `512Mi`             |
| `extraVolumes`                       | Optionally specify extra list of additional volumes for the ZooKeeper pod(s)                                                                                                                     | `[]`                |
| `extraVolumeMounts`                  | Optionally specify extra list of additional volumeMounts for the ZooKeeper container(s)                                                                                                          | `[]`                |
| `extraEnvVars`                       | Array with extra environment variables to add to ZooKeeper nodes                                                                                                                                 | `[]`                |


### Traffic Exposure parameters

| Name                                        | Description                                                                             | Value       |
| ------------------------------------------- | --------------------------------------------------------------------------------------- | ----------- |
| `service.type`                              | Kubernetes Service type                                                                 | `ClusterIP` |
| `service.ports.client`                      | ZooKeeper client service port                                                           | `2181`      |
| `service.ports.tls`                         | ZooKeeper TLS service port                                                              | `3181`      |
| `service.ports.follower`                    | ZooKeeper follower service port                                                         | `2888`      |
| `service.ports.election`                    | ZooKeeper election service port                                                         | `3888`      |
| `service.nodePorts.client`                  | Node port for clients                                                                   | `""`        |
| `service.nodePorts.tls`                     | Node port for TLS                                                                       | `""`        |
| `service.disableBaseClientPort`             | Remove client port from service definitions.                                            | `false`     |
| `service.clusterIP`                         | ZooKeeper service Cluster IP                                                            | `""`        |
| `service.loadBalancerIP`                    | ZooKeeper service Load Balancer IP                                                      | `""`        |
| `service.loadBalancerSourceRanges`          | ZooKeeper service Load Balancer sources                                                 | `[]`        |
| `service.annotations`                       | Additional custom annotations for ZooKeeper service                                     | `{}`        |
| `service.headless.annotations`              | Annotations for the Headless Service                                                    | `{}`        |
| `service.headless.publishNotReadyAddresses` | If the ZooKeeper headless service should publish DNS records for not ready pods         | `true`      |


### Other Parameters

| Name                                          | Description                                                            | Value   |
| --------------------------------------------- | ---------------------------------------------------------------------- | ------- |
| `serviceAccount.create`                       | Enable creation of ServiceAccount for ZooKeeper pod                    | `false` |
| `serviceAccount.name`                         | The name of the ServiceAccount to use.                                 | `""`    |
| `serviceAccount.automountServiceAccountToken` | Allows auto mount of ServiceAccountToken on the serviceAccount created | `true`  |
| `serviceAccount.annotations`                  | Additional custom annotations for the ServiceAccount                   | `{}`    |


### Persistence parameters

| Name                                | Description                                                                    | Value               |
|-------------------------------------| ------------------------------------------------------------------------------ | ------------------- |
| `persistence.enabled`               | Enable ZooKeeper data persistence using PVC. If false, use emptyDir            | `true`              |
| `persistence.storageClassName`      | PVC Storage Class for ZooKeeper data volume                                    | `""`                |
| `persistence.accessModes`           | PVC Access modes                                                               | `["ReadWriteOnce"]` |
| `persistence.annotations`           | Annotations for the PVC                                                        | `{}`                |
| `persistence.data.size`             | PVC Storage Request for ZooKeeper data volume                                  | `8Gi`               |
| `persistence.data.selector`         | Selector to match an existing Persistent Volume for ZooKeeper's data PVC       | `{}`                |
| `persistence.data.existingClaim`    | Name of an existing PVC to use (only when deploying a single replica)          | `""`                |
| `persistence.dataLog.size`          | PVC Storage Request for ZooKeeper's dedicated data log directory               | `8Gi`               |
| `persistence.dataLog.selector`      | Selector to match an existing Persistent Volume for ZooKeeper's data log PVC   | `{}`                |
| `persistence.dataLog.existingClaim` | Provide an existing `PersistentVolumeClaim` for ZooKeeper's data log directory | `""`                |


### Metrics parameters

| Name                                       | Description                                                                           | Value       |
| ------------------------------------------ | ------------------------------------------------------------------------------------- | ----------- |
| `metrics.enabled`                          | Enable Prometheus to access ZooKeeper metrics endpoint                                | `false`     |
| `metrics.containerPort`                    | ZooKeeper Prometheus Exporter container port                                          | `9141`      |
| `metrics.service.type`                     | ZooKeeper Prometheus Exporter service type                                            | `ClusterIP` |
| `metrics.service.port`                     | ZooKeeper Prometheus Exporter service port                                            | `9141`      |
| `metrics.service.annotations`              | Annotations for Prometheus to auto-discover the metrics endpoint                      | `{}`        |
| `metrics.serviceMonitor.enabled`           | Create ServiceMonitor Resource for scraping metrics using Prometheus Operator         | `false`     |
| `metrics.serviceMonitor.namespace`         | Namespace for the ServiceMonitor Resource (defaults to the Release Namespace)         | `""`        |
| `metrics.serviceMonitor.interval`          | Interval at which metrics should be scraped.                                          | `""`        |
| `metrics.serviceMonitor.scrapeTimeout`     | Timeout after which the scrape is ended                                               | `""`        |
| `metrics.serviceMonitor.additionalLabels`  | Additional labels that can be used so ServiceMonitor will be discovered by Prometheus | `{}`        |
| `metrics.serviceMonitor.selector`          | Prometheus instance selector labels                                                   | `{}`        |
| `metrics.serviceMonitor.relabelings`       | RelabelConfigs to apply to samples before scraping                                    | `[]`        |
| `metrics.serviceMonitor.metricRelabelings` | MetricRelabelConfigs to apply to samples before ingestion                             | `[]`        |
| `metrics.serviceMonitor.honorLabels`       | Specify honorLabels parameter to add the scrape endpoint                              | `false`     |
| `metrics.serviceMonitor.jobLabel`          | The name of the label on the target service to use as the job name in prometheus.     | `""`        |
| `metrics.prometheusRule.enabled`           | Create a PrometheusRule for Prometheus Operator                                       | `false`     |
| `metrics.prometheusRule.namespace`         | Namespace for the PrometheusRule Resource (defaults to the Release Namespace)         | `""`        |
| `metrics.prometheusRule.additionalLabels`  | Additional labels that can be used so PrometheusRule will be discovered by Prometheus | `{}`        |
| `metrics.prometheusRule.rules`             | PrometheusRule definitions                                                            | `[]`        |


### TLS/SSL parameters

| Name                                      | Description                                                                                        | Value                                                                 |
| ----------------------------------------- | -------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `tls.client.enabled`                      | Enable TLS for client connections                                                                  | `false`                                                               |
| `tls.client.auth`                         | SSL Client auth. Can be "none", "want" or "need".                                                  | `none`                                                                |
| `tls.client.autoGenerated`                | Generate automatically self-signed TLS certificates for ZooKeeper client communications            | `false`                                                               |
| `tls.client.existingSecret`               | Name of the existing secret containing the TLS certificates for ZooKeeper client communications    | `""`                                                                  |
| `tls.client.existingSecretKeystoreKey`    | The secret key from the tls.client.existingSecret containing the Keystore.                         | `""`                                                                  |
| `tls.client.existingSecretTruststoreKey`  | The secret key from the tls.client.existingSecret containing the Truststore.                       | `""`                                                                  |
| `tls.client.keystorePath`                 | Location of the KeyStore file used for Client connections                                          | `/opt/bitnami/zookeeper/config/certs/client/zookeeper.keystore.jks`   |
| `tls.client.truststorePath`               | Location of the TrustStore file used for Client connections                                        | `/opt/bitnami/zookeeper/config/certs/client/zookeeper.truststore.jks` |
| `tls.client.passwordsSecretName`          | Existing secret containing Keystore and truststore passwords                                       | `""`                                                                  |
| `tls.client.passwordsSecretKeystoreKey`   | The secret key from the tls.client.passwordsSecretName containing the password for the Keystore.   | `""`                                                                  |
| `tls.client.passwordsSecretTruststoreKey` | The secret key from the tls.client.passwordsSecretName containing the password for the Truststore. | `""`                                                                  |
| `tls.client.keystorePassword`             | Password to access KeyStore if needed                                                              | `""`                                                                  |
| `tls.client.truststorePassword`           | Password to access TrustStore if needed                                                            | `""`                                                                  |
| `tls.quorum.enabled`                      | Enable TLS for quorum protocol                                                                     | `false`                                                               |
| `tls.quorum.auth`                         | SSL Quorum Client auth. Can be "none", "want" or "need".                                           | `none`                                                                |
| `tls.quorum.autoGenerated`                | Create self-signed TLS certificates. Currently only supports PEM certificates.                     | `false`                                                               |
| `tls.quorum.existingSecret`               | Name of the existing secret containing the TLS certificates for ZooKeeper quorum protocol          | `""`                                                                  |
| `tls.quorum.existingSecretKeystoreKey`    | The secret key from the tls.quorum.existingSecret containing the Keystore.                         | `""`                                                                  |
| `tls.quorum.existingSecretTruststoreKey`  | The secret key from the tls.quorum.existingSecret containing the Truststore.                       | `""`                                                                  |
| `tls.quorum.keystorePath`                 | Location of the KeyStore file used for Quorum protocol                                             | `/opt/bitnami/zookeeper/config/certs/quorum/zookeeper.keystore.jks`   |
| `tls.quorum.truststorePath`               | Location of the TrustStore file used for Quorum protocol                                           | `/opt/bitnami/zookeeper/config/certs/quorum/zookeeper.truststore.jks` |
| `tls.quorum.passwordsSecretName`          | Existing secret containing Keystore and truststore passwords                                       | `""`                                                                  |
| `tls.quorum.passwordsSecretKeystoreKey`   | The secret key from the tls.quorum.passwordsSecretName containing the password for the Keystore.   | `""`                                                                  |
| `tls.quorum.passwordsSecretTruststoreKey` | The secret key from the tls.quorum.passwordsSecretName containing the password for the Truststore. | `""`                                                                  |
| `tls.quorum.keystorePassword`             | Password to access KeyStore if needed                                                              | `""`                                                                  |
| `tls.quorum.truststorePassword`           | Password to access TrustStore if needed                                                            | `""`                                                                  |
| `tls.resources.limits`                    | The resources limits for the TLS init container                                                    | `{}`                                                                  |
| `tls.resources.requests`                  | The requested resources for the TLS init container                                                 | `{}`                                                                  |

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

