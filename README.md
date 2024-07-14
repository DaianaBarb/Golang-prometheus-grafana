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
- Vá para Dashboards clique em + para adicionar o dashboard e cole o json que esta dentro da pasta grafana neste projeto.
- irá aparecer o dashboard com as metricas adicionadas.
- Se quiser adicionar mais metricas ao projeto e so clicar em editar  nos tres pontinhos que ficam do lado direito em cima do dashboard e clicar em add query.
