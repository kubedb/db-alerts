{{/*
Expand the name of the chart.
*/}}
{{- define "redis.name" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "redis.fullname" -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "redis.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "redis.labels" -}}
helm.sh/chart: {{ include "redis.chart" . }}
{{ include "redis.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "redis.selectorLabels" -}}
app.kubernetes.io/name: {{ include "redis.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Alert labels
*/}}
{{- define "redis.alertLabels" -}}
k8s-group: {{ .Values.metadata.resource.group }}
k8s-kind: {{ .Values.metadata.resource.kind }}
k8s-resource: {{ .Values.metadata.resource.name }}
k8s-name: {{ include "redis.fullname" . }}
k8s-namespace: {{ .Release.Namespace }}
{{- if .Values.spec.alert.additionalRuleLabels }}
{{- toYaml .Values.spec.alert.additionalRuleLabels }}
{{- end }}
{{- end }}
