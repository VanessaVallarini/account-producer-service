import os, sys
from urllib.request import urlretrieve

os.chdir(os.path.dirname(sys.argv[0]))

from diagrams import Cluster, Diagram, Edge
from diagrams.k8s.compute import Pod
from diagrams.onprem.queue import Kafka
from diagrams.custom import Custom
from diagrams.programming.language import Go
from diagrams.k8s.network import Service

with Diagram("account producer service", show = False, direction="TB"):
    blueline=Edge(color="blue",style="bold")
    darkOrange=Edge(color="darkOrange",style="bold")
    blackline=Edge(color="black",style="bold")
    scylladb_url = "https://upload.wikimedia.org/wikipedia/en/2/20/Scylla_the_sea_monster.png"
    scylladb_icon = "scylla.png"
    urlretrieve(scylladb_url, scylladb_icon)

    with Cluster("account-producer-pod"):
        producerPod=Pod("account-producer-pod")

    with Cluster("external"):
       producerCreateKafka=Kafka("account_createorupdate")
       producerUpdateKafka=Kafka("account_delete") 
       producerViaCepApi=Service("via-cep-api")

    with Cluster("internal"):
       accountAvros=Go("account-toolkit")

    with Cluster("scyllaDb"):
       accountDatabase=Custom("account-database",scylladb_icon)

    producerPod - darkOrange >> producerCreateKafka
    producerPod - darkOrange >> producerUpdateKafka
    producerPod - blackline >> accountAvros
    producerPod - blackline >> producerViaCepApi
    producerPod - blueline >> accountDatabase