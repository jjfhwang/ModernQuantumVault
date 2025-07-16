# ModernQuantumVault: Secure, Quantum-Resistant Key Management

ModernQuantumVault is a Go-based key management system designed with a focus on security and future-proofing against potential quantum computing threats. It provides a robust and flexible solution for generating, storing, and managing cryptographic keys, secrets, and other sensitive data, ensuring the confidentiality and integrity of your critical assets. Unlike traditional key vaults, ModernQuantumVault incorporates post-quantum cryptography (PQC) algorithms alongside classical cryptography, offering a layered defense against both current and emerging threats. This allows organizations to prepare for the future while maintaining compatibility with existing infrastructure.

The primary objective of ModernQuantumVault is to simplify the complexities of key management in modern distributed systems. It addresses the challenges of securely storing and accessing keys across multiple applications and environments by providing a centralized, auditable, and policy-driven key management service. The system supports various encryption algorithms, including AES, ChaCha20, and PQC candidates like Kyber and Dilithium, allowing users to choose the algorithms that best suit their specific security requirements and performance constraints. Furthermore, it provides fine-grained access control mechanisms to ensure that only authorized users and applications can access sensitive data.

ModernQuantumVault goes beyond simple key storage by offering advanced features such as key rotation, versioning, and auditing. Key rotation automates the process of periodically generating new keys and retiring old ones, reducing the risk of key compromise. Versioning allows users to track the history of keys and revert to previous versions if necessary. Auditing provides a detailed log of all key management operations, enabling organizations to monitor activity, detect potential security breaches, and comply with regulatory requirements. The entire architecture is designed with resilience and scalability in mind, ensuring that the system can handle the demands of modern enterprise environments.

## Key Features

*   **Quantum-Resistant Cryptography:** Implements post-quantum cryptographic algorithms (e.g., Kyber, Dilithium) alongside traditional algorithms to mitigate future quantum computing risks. Algorithm choices are configurable per key.
*   **Hierarchical Key Management:** Supports a hierarchical key structure, enabling organizations to organize keys according to their specific needs and apply granular access controls at each level. This allows for logical separation of concerns and simplifies key management in complex environments.
*   **Role-Based Access Control (RBAC):** Enforces strict access control policies based on user roles and permissions. Roles can be defined with specific privileges to access, manage, and use keys. Access control lists (ACLs) are defined using a Go-based policy language, providing flexibility and expressiveness.
*   **Key Rotation and Versioning:** Automates key rotation and maintains a complete history of all key versions. This reduces the risk of key compromise and provides the ability to revert to previous versions if necessary. The rotation interval and method (e.g., manual, scheduled) are configurable.
*   **Audit Logging:** Provides a comprehensive audit trail of all key management operations, including key creation, deletion, access, and modification. Logs are stored in a tamper-proof format and can be integrated with security information and event management (SIEM) systems. Log retention policies are configurable.
*   **API-Driven Interface:** Offers a well-defined REST API for programmatic access to all key management functions. The API is documented using OpenAPI/Swagger specifications, making it easy for developers to integrate ModernQuantumVault into their applications.
*   **Hardware Security Module (HSM) Integration:** Supports integration with HSMs for secure key storage and cryptographic operations. This provides an additional layer of security by protecting keys from software-based attacks. Integration is achieved through the PKCS#11 interface.

## Technology Stack

*   **Go:** The primary programming language used for building the core key management functionality. Go was chosen for its performance, concurrency support, and strong standard library.
*   **gRPC:** Used for internal communication between services within ModernQuantumVault. gRPC provides a high-performance, efficient, and reliable communication protocol.
*   **Protocol Buffers:** Used for defining the data structures and service interfaces for gRPC communication. Protocol Buffers provide a language-neutral, platform-neutral, extensible mechanism for serializing structured data.
*   **etcd:** A distributed key-value store used for storing configuration data, service discovery, and leader election. etcd provides a reliable and consistent storage backend for critical metadata.
*   **PostgreSQL:** Used for storing key metadata, audit logs, and access control policies. PostgreSQL provides a robust and scalable relational database system.

## Installation

1.  Ensure you have Go 1.20 or later installed.
2.  Clone the repository:
    git clone https://github.com/jjfhwang/ModernQuantumVault.git
3.  Navigate to the project directory:
    cd ModernQuantumVault
4.  Install dependencies:
    go mod download
5.  Build the application:
    go build -o modernquantumvault cmd/modernquantumvault/main.go

## Configuration

ModernQuantumVault is configured using environment variables. The following variables are required:

*   `MQV_DATABASE_URL`: The connection string for the PostgreSQL database. Example: `postgres://user:password@host:port/database`
*   `MQV_ETCD_ENDPOINTS`: A comma-separated list of etcd endpoints. Example: `http://127.0.0.1:2379,http://127.0.0.1:2380`
*   `MQV_LISTEN_ADDRESS`: The address on which the API server will listen. Example: `:8080`
*   `MQV_HSM_LIBRARY`: The path to the HSM library (optional). Required for HSM integration.
*   `MQV_HSM_PIN`: The PIN for accessing the HSM (optional). Required for HSM integration.

You can set these environment variables in your shell or create a `.env` file in the project directory.

## Usage

After installing and configuring ModernQuantumVault, you can start the API server:

./modernquantumvault

The API is accessible via HTTP at the address specified by the `MQV_LISTEN_ADDRESS` environment variable. A Swagger/OpenAPI definition of the API is available at `/swagger/index.html`.

Example API calls (using curl):

* Create a new key:

* Get a key:

* Encrypt data:

* Decrypt data:


Detailed API documentation, including request and response schemas, is available in the Swagger/OpenAPI definition.

## Contributing

We welcome contributions to ModernQuantumVault. Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write tests for your code.
4.  Ensure all tests pass before submitting a pull request.
5.  Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/ModernQuantumVault/blob/main/LICENSE) file for details.

## Acknowledgements

This project was inspired by the need for secure and quantum-resistant key management solutions in modern distributed systems. We would like to thank the developers of the underlying technologies, including Go, gRPC, etcd, and PostgreSQL, for their contributions to the open-source community.