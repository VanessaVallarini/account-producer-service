#!/bin/bash

WAITING_TIME_SEC=3

CONFIG_DIR=/home/cql;
if [ $# -gt 0 ]; then
    CONFIG_DIR=$1;
fi

if [ -z "${SCYLLA_USER}" ]; then
    SCYLLA_USER=cassandra
fi
if [ -z "${SCYLLA_PASSWORD}" ]; then
    SCYLLA_PASSWORD=cassandra
fi

echo "Scylla host: ${CQLSH_HOST}"

# wait for default user creation
while true ; do
    echo "Waiting node initialization..."
    cqlsh -u ${SCYLLA_USER} -p ${SCYLLA_PASSWORD} -e "SHOW HOST;"
    RESULT=$?

    [[ "${RESULT}" != "0" ]] || break
    sleep ${WAITING_TIME_SEC}
done
echo "Node initialized. Applying cql files..."

echo "CQL files dir: ${CONFIG_DIR}"

# execute every cql file
for FILE in ${CONFIG_DIR}/*.cql; do
    echo "Applying file: ${FILE}"
    cqlsh -u ${SCYLLA_USER} -p ${SCYLLA_PASSWORD} -f ${FILE}
    if [ $? != 0 ] ; then
        echo "Erou."
    fi
done

echo "All .cql files applied"