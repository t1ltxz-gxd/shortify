#!/bin/bash

# Create a .env file if it doesn't exist
if [ ! -f .env ]; then
    touch .env
fi

# Function to ask for input and write to .env
write_env() {
    local var_name=$1
    local prompt_message=$2
    local default_value=$3
    local secret=$4

    while true; do
        if [ "$secret" == "true" ]; then
            read -rsp "Enter the $prompt_message (default: $default_value): " input
            echo
        else
            read -rp "Enter the $prompt_message (default: $default_value): " input
        fi

        # Use the default value if input is empty
        input=${input:-$default_value}

        # Validate ENV value
        if [ "$var_name" == "ENV" ]; then
            if [[ "$input" =~ ^(dev|development|prod|production)$ ]]; then
                break
            else
                echo "Invalid value for ENV. Allowed values are dev, development, prod, production."
                continue
            fi
        fi

        # Validate POSTGRES_PORT value
        if [ "$var_name" == "POSTGRES_PORT" -o "$var_name" == "GRPC_PORT" ]; then
            if [[ "$input" =~ ^[0-9]+$ ]] && [ "$input" -ge 1 ] && [ "$input" -le 65535 ]; then
                break
            else
                echo "Invalid value for POSTGRES_PORT. It should be an integer between 1 and 65535."
                continue
            fi
        fi

        # Validate REDIS_PASS and POSTGRES_PASSWORD length
        if [ "$var_name" == "REDIS_PASS" -o "$var_name" == "POSTGRES_PASSWORD" ]; then
            if [ ${#input} -le 128 ]; then
                break
            else
                echo "Invalid value for $var_name. It should be no more than 128 characters."
                continue
            fi
        fi

        break
    done

    # Check if the variable already exists in the .env file
    if grep -q "^$var_name=" .env; then
        # If the variable exists, replace it with the new value
        sed -i "/^$var_name=/c\\$var_name=$input" .env
    else
        # If the variable doesn't exist, append it to the .env file
        echo "$var_name=$input" >> .env
    fi
}

# Ask for and write environment variables
write_env "ENV" "environment" "development" "false"
write_env "GRPC_PORT" "gRPC server port" "50051" "false"
write_env "REDIS_PASS" "Redis password" "mysecretpassword" "true"
write_env "POSTGRES_USER" "PostgreSQL user" "postgres" "false"
write_env "POSTGRES_PASSWORD" "PostgreSQL password" "postgres" "true"
write_env "POSTGRES_PORT" "PostgreSQL port" "5432" "false"
write_env "POSTGRES_DB" "PostgreSQL database name" "postgres" "false"

# Display a message to the user
echo ".env file with variables created successfully."