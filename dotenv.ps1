# Check if the .env file exists. If not, create it.
if (!(Test-Path -Path .env)) {
    New-Item -ItemType File -Path .env -Force
}

# Function to ask for input and write to .env
function Write-Env {
    param (
        [string]$varName,
        [string]$promptMessage,
        [string]$defaultValue,
        [bool]$secret
    )

    do {
        if ($secret) {
            $input = Read-Host -Prompt "Enter the $promptMessage (default: $defaultValue)" -AsSecureString
            $input = [Runtime.InteropServices.Marshal]::PtrToStringAuto([Runtime.InteropServices.Marshal]::SecureStringToBSTR($input))
        } else {
            $input = Read-Host -Prompt "Enter the $promptMessage (default: $defaultValue)"
        }

        # Use the default value if input is empty
        if ([string]::IsNullOrEmpty($input)) {
            $input = $defaultValue
        }

        # Validate ENV value
        if ($varName -eq "ENV") {
            if ($input -match "^(dev|development|prod|production)$") {
                break
            } else {
                Write-Host "Invalid value for ENV. Allowed values are dev, development, prod, production."
                continue
            }
        }

        # Validate POSTGRES_PORT value
        if ($varName -eq "POSTGRES_PORT" -or $varName -eq "GRPC_PORT") {
            if ($input -match "^[0-9]+$" -and $input -ge 1 -and $input -le 65535) {
                break
            } else {
                Write-Host "Invalid value for POSTGRES_PORT. It should be an integer between 1 and 65535."
                continue
            }
        }

        # Validate REDIS_PASS and POSTGRES_PASSWORD length
        if ($varName -eq "REDIS_PASS" -or $varName -eq "POSTGRES_PASSWORD") {
            if ($input.Length -le 128) {
                break
            } else {
                Write-Host "Invalid value for $varName. It should be no more than 128 characters."
                continue
            }
        }

        break
    } while ($true)

    # Check if the variable already exists in the .env file
    $content = Get-Content -Path .env
    if ($content -match "^$varName=") {
        # If the variable exists, replace it with the new value
        $content = $content -replace "^$varName=.*", "$varName=$input"
        $content | Set-Content -Path .env
    } else {
        # If the variable doesn't exist, append it to the .env file
        Add-Content -Path .env -Value "$varName=$input"
    }
}

# Ask for and write environment variables
Write-Env -varName "ENV" -promptMessage "environment" -defaultValue "development" -secret $false
Wrive-Env -varName "GRPC_PORT" -promptMessage "gRPC port" -defaultValue "50051" -secret $false
Write-Env -varName "REDIS_PASS" -promptMessage "Redis password" -defaultValue "mysecretpassword" -secret $true
Write-Env -varName "POSTGRES_USER" -promptMessage "PostgreSQL user" -defaultValue "postgres" -secret $false
Write-Env -varName "POSTGRES_PASSWORD" -promptMessage "PostgreSQL password" -defaultValue "postgres" -secret $true
Write-Env -varName "POSTGRES_PORT" -promptMessage "PostgreSQL port" -defaultValue "5432" -secret $false
Write-Env -varName "POSTGRES_DB" -promptMessage "PostgreSQL database name" -defaultValue "postgres" -secret $false

# Display a message to the user
Write-Host ".env file with variables created successfully."