FROM alpine/helm:3.7.1

RUN helm plugin install  https://github.com/chartmuseum/helm-push

RUN helm repo add westwell https://wellharbor.westwell-research.com/chartrepo --username=admin --password=Westwell@2020

RUN helm repo add spiking https://wellharbor.westwell-research.com/chartrepo/well_spiking --username=admin --password=Westwell@2020