#!/bin/bash

deploy_chart() {
    helm upgrade dynamic-pv-scaler ./deploy/helm -f \
    ./deploy/helm/values.yaml --install --namespace dynamic-pv-scaler
}

validate_chart() {
    helm test dynamic-pv-scaler --namespace dynamic-pv-scaler
}

validate_container_state() {
    output=$(kubectl get pods -n keycloak -l app=dynamic-pv-scaler \
    -o jsonpath="{.items[*]['status.phase']}")
    if [ "${output}" != "Running" && "${output}" != "" ]
    then
        echo "Container is not healthy"
        exit 1
    else
        echo "Container is running fine"
    fi
}

main_function() {
    deploy_chart
    validate_chart
    sleep 30s
    validate_container_state
}

main_function