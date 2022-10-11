package make

import (
	"bytes"
	"dtweave.io/zookeeper-operator/api/v1alpha1"
	"fmt"
	"html/template"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func Configmap(instance *v1alpha1.Zookeeper) (*v1.ConfigMap, error) {
	cfg, err := parseTemplate(&instance.Spec.Conf)
	if err != nil {
		return nil, err
	}

	if instance.Spec.Conf.AdditionalConfig != nil {
		for k, v := range instance.Spec.Conf.AdditionalConfig {
			cfg = cfg + fmt.Sprintf("%s=%s\n", k, v)
		}
	}

	cm := &v1.ConfigMap{
		ObjectMeta: v12.ObjectMeta{
			Name:      instance.ConfigMapName(),
			Namespace: instance.Namespace,
		},
		Data: map[string]string{
			"zoo.cfg": cfg,
		},
	}
	if instance.Spec.CommonAnnotations != nil {
		cm.ObjectMeta.Annotations = instance.Spec.CommonAnnotations
	}
	if instance.Spec.CommonLabels != nil {
		cm.ObjectMeta.Labels = instance.Spec.CommonLabels
	}

	return cm, nil
}

func parseTemplate(conf *v1alpha1.ZookeeperConf) (string, error) {
	path, err2 := os.Getwd()
	if err2 != nil {
		fmt.Println("unable to get dir")
	}

	tmpl, err := template.ParseFiles(path + "/pkg/template/zoo.cfg")
	if err != nil {
		return "", err
	}

	b := new(bytes.Buffer)
	err = tmpl.Execute(b, conf)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
