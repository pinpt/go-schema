-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

{{- $restn := .ResourceTable }}
{{- $rescn := .ResourceColumns }}
{{- $idcol := .IDColumn }}
{{- $roletn := .RoleTable }}
{{- $rolecn := .RoleColumns }}
{{ range .Routes }}
INSERT INTO {{ $restn }} ({{ $rescn }}) VALUES ({{ .SQLValues }});
{{- if .Auth -}}
{{ range $k, $f := .Auth.Features }}
INSERT INTO {{ $restn }} ({{ $rescn }}) VALUES ({{ $f.SQLValues }});
{{- end }}
{{- end }}
{{- end}}

{{- range .Roles }}
INSERT INTO {{ $roletn }} ({{ $rolecn }}) VALUES ({{ .SQLValues }});
{{- end}}

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

{{ range .Routes }}
DELETE FROM {{ $restn }} WHERE {{ $idcol }} = {{ .SQLID }};
{{- if .Auth -}}
{{ range $k, $f := .Auth.Features }}
DELETE FROM {{ $restn }} WHERE {{ $idcol }} = {{ .SQLID }};
{{- end }}
{{- end }}
{{- end}}

{{- range .Roles }}
DELETE FROM {{ $roletn }} WHERE {{ $idcol }} = {{ .SQLID }};
{{- end}}
