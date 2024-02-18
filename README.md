<img alt="LSR Cloud" width="40%" src="https://lucianoromao.com.br/lsr.png">

# Estudo sobre Sistemas Baseados em Micro Serviços

:heavy_exclamation_mark:
Esse repositório contém apenas um microserviço, sendo ele o WallerCore.

Nesse link você poderá encontrar todos os microserviços se comunicando e já prontos para execução com docker: TBD
:heavy_exclamation_mark:

## Visão Geral 🔎
Neste código espero mostrar a minha visão de sistemas baseados em Micro Serviços utilizando PlantUML para a criação diagramas usando o módelo C4.

## O que é um sistema baseado em Micro Serviços? :thinking:
 * São aplicações comuns, ou seja, não é um bicho de sete cabeças.
 * É um "grande" sistema, que é feito de diversos "micro sistemas".
 * Esses micro serviço devem ter objetivos bem definidos e únicos.
 * Os micros serviços devem sempre ser autônomos.
 * Cada micro serviço tem seu próprio deploy, ou seja, seu próprio CI/CD.

## Reflexões sobre micro serviços :factory:
Sistemas baseados em Micro Serviços ajudam muito quando temos um sistemas complexo e que necessita de alta disponibilidade, pois nos possibilita esacalar partes espefícas dos sistemas conforme necessidade, sem disperdiçar recursos. Entretando existe uma grande complexidade em desenvolver e sustentar esse tipo de sistema. Sendo assim não é recomendado que um sistema já nasça com um arquitetura baseada em micro serviços, o ideal é o que o sistema tenha sido primeiramente criado em uma versão monolito modular e primeiro passe por uma prova de conceito.
    
Se pensarmos na parte de sustenção de um sistema baseado em micro serviços o primeiro ponto que devemos pensar é sobre um padrão os quais esses serviços devem seguir como, por exemplo, esteiras de CI/CD, linguagem de programação, banco de dados e etc. Eu considereo essa parte como a base do projeto pois isso irá definir o qual dificil será manter e encontrar profissionais qualificados para trabalhar em um sistema desse tipo.

Com uma base solida vamos começar a penasr nas futuras complexidades de um sistema baseado em micro serviços, dado que esse tipo cada micro serviço precisa ter seu próprio objetivo podemos chegar em um ponto em que isso se torna "caótico" no ponto de vista de gerenciamento devido a grande quantidade de micro serviços que possa vir a existir. Para ajudar a sustentar e manter o projeto vivo existem alguns padrões que podem ser seguidos.

 * (API Composition)[https://microservices.io/patterns/data/api-composition.html]
 * (Decompose by business capability)[https://microservices.io/patterns/decomposition/decompose-by-business-capability.html]
 * (Strangler Application)[https://microservices.io/patterns/refactoring/strangler-application.html]
 * (ACL)[https://microservices.io/patterns/refactoring/anti-corruption-layer.html]
 * (API Gateway)[https://microservices.io/patterns/apigateway.html]
 * (BFF)[https://medium.com/digitalproductsdev/arquitetura-bff-back-end-for-front-end-13e2cbfbcda2]
 * (Transactional Outbox)[https://microservices.io/patterns/data/transactional-outbox.html]

## Abordagem :construction_worker:
Irei utilizar um sistema simples, mas que irá representar bem a essência dos micro serviços, e para isso irei documentar o projeto usandoi diagramas do modelo (C4)[https://c4model.com/]
