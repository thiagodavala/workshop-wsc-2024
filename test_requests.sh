#!/bin/bash
# Definir o número de requisições
NUM_REQUESTS=10
CONCURRENT_REQUESTS=1  # Número de requisições simultâneas

# Testar a rota /health
echo "Iniciando requisições para /health..."
ab -n $NUM_REQUESTS -c $CONCURRENT_REQUESTS "$BASE_URL/health"
echo "Requisições para /health finalizadas."
echo ""

# Testar a rota /calc?input=ddfdf
echo "Iniciando requisições para /calc?input=ddfdf..."
ab -n $NUM_REQUESTS -c $CONCURRENT_REQUESTS "$BASE_URL/calc?input=ddfdf"
echo "Requisições para /calc?input=ddfdf finalizadas."
