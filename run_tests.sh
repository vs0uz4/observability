#!/bin/bash

# Garantir que o script use o locale padrão com ponto como separador decimal
export LC_NUMERIC="C"

services=("ms-weatherzip" "ms-searchzip")
overall_coverage=0

echo "Starting tests for all services..."

for service in "${services[@]}"; do
  echo "---------------------------------------------------------------------------------------------------------------------"
  echo "Running tests for $service..."
  if (cd $service && go test -timeout 30s -v ./... -coverprofile=report/coverage.out); then
    echo "Calculating coverage for $service..."
    service_coverage=$(cd $service && go tool cover -func=report/coverage.out | grep total | awk '{print $3}' | sed 's/%//')
    echo "$service Coverage: $service_coverage%"
    overall_coverage=$(echo "$overall_coverage + $service_coverage" | bc)
  else
    echo "Tests failed for $service."
  fi
done

echo "---------------------------------------------------------------------------------------------------------------------"
echo "All tests completed."

num_services=${#services[@]}
average_coverage=$(echo "scale=2; $overall_coverage / $num_services" | bc)
printf "Overall Coverage: %.2f%%\n" "$average_coverage"
