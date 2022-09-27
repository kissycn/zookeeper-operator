package make

import (
	"bytes"
	"dtweave.io/zookeeper-operator/api/v1alpha1"
	"html/template"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Configmap(instance *v1alpha1.Zookeeper) (*v1.ConfigMap, error) {
	cfg, err := parseTemplate(&instance.Spec.Conf)
	if err != nil {
		return nil, err
	}

	return &v1.ConfigMap{
		//TypeMeta: v12.TypeMeta{
		//	Kind:       "ConfigMap",
		//	APIVersion: "v1",
		//},
		ObjectMeta: v12.ObjectMeta{
			Name:        instance.Name,
			Namespace:   instance.Namespace,
			Annotations: instance.Annotations,
			Labels:      instance.Labels,
		},
		Data: map[string]string{
			"zoo.cfg": cfg,
		},
	}, nil
}

func parseTemplate(conf *v1alpha1.ZookeeperConf) (string, error) {
	tmpl, err := template.ParseFiles("../template/zoo.cfg")
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
