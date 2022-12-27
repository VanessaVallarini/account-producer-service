import os, sys

os.chdir(os.path.dirname(sys.argv[0]))

from diagrams import Cluster, Diagram, Edge
from diagrams.k8s.compute import Pod
from diagrams.onprem.queue import Kafka
from diagrams.aws.network import Route53


with Diagram("account producer service", show = False):
    blueline=Edge(color="blue",style="bold")
    darkOrange=Edge(color="darkOrange",style="bold")
    blackline=Edge(color="black",style="bold")

    with Cluster("account-endpoint"):
        accountEndpoint=Route53("account-endpoint")
    
    with Cluster("account-producer-pod"):
        producerPod=Pod("producer-pod")

    with Cluster("kafka-producer"):
        producerKafka=Kafka("producer-kafka")

    accountEndpoint - blueline >> producerPod
    producerPod - darkOrange >> producerKafka