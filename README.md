Projetinho simples so para ver o funcionamento do prometheus + grafana com golang.

Nao usei boas praticas e nem organizei as pastas pq e so um exemplo simples.

Primeiramente rode o dockerfile com o terminal na raiz do projeto para gerar a imagem do Projeto com o comando:

- docker build -t image-golang-test .


Rode o docker compose com o comando:
- docker-compose up 

Acesse o grafana na porta localhost:3000/login

O login/senha e: admin/admin

Para conectar o grafana com o prometheus:

- Dentro do grafana vá em Connections e clique em add new Connection.
- Procure por Prometheus.
- Clique em add new 
- Adicione a porta do prometheus que definimos no docker compose que e a porta: http://prometheus:9090
- Marque a opção Default
- Clique em save & test
- Aparecerá a mensagem: "Successfully queried the Prometheus API.
Next, you can start to visualize data by building a dashboard, or by querying data in the Explore view."
- Vá para Dashboards clique em + para adicionar o dashboard e cole este json:
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "id": 1,
  "links": [],
  "panels": [
    {
      "datasource": null,
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 6,
      "options": {
        "displayMode": "gradient",
        "orientation": "horizontal",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "showUnfilled": true,
        "text": {}
      },
      "pluginVersion": "8.0.6",
      "targets": [
        {
          "exemplar": true,
          "expr": "http_request_get_user_status_count",
          "interval": "",
          "legendFormat": "{{status}}  {{user}}",
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "Panel Title",
      "type": "bargauge"
    }
  ],
  "refresh": "",
  "schemaVersion": 30,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "user_status",
  "uid": "g5kaN3W7k",
  "version": 2
}
- irá aparecer o dashboard com as metricas adicionadas.
- Se quiser adicionar mais metricas ao projeto e so clicar em editar  nos tres pontinhos que ficam do lado direito em cima do dashboard e clicar em add query.
