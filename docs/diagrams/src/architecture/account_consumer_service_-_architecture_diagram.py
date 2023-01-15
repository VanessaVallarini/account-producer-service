import os, sys

os.chdir(os.path.dirname(sys.argv[0]))

from diagrams import Cluster, Diagram, Edge
from diagrams.k8s.compute import Pod
from diagrams.onprem.queue import Kafka
from diagrams.aws.database import Aurora
from diagrams.k8s.network import Service


with Diagram("account producer service", show = False):
    blueline=Edge(color="blue",style="bold")
    darkOrange=Edge(color="darkOrange",style="bold")
    blackline=Edge(color="black",style="bold")

    with Cluster("account-producer-pod"):
        producerPod=Pod("account-producer-pod")

    with Cluster("external"):
       producerCreateKafka=Kafka("account-create")
       producerUpdateKafka=Kafka("account-update") 
       accountAvros=Service("account-toolkit")  
       producerViaCepApi=Service("via-cep-api")

    with Cluster("scyllaDb"):
       accountDatabase=Aurora("account-database")

    producerPod - darkOrange >> producerCreateKafka
    producerPod - darkOrange >> producerUpdateKafka
    producerPod - blackline >> accountAvros
    producerPod - blackline >> producerViaCepApi
    producerPod - blueline >> accountDatabase