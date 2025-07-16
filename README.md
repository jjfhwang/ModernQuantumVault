# ModernQuantumVault: Decentralized Crypto-Asset Data Aggregation with Tamper-Evident History

ModernQuantumVault provides a decentralized and secure solution for aggregating crypto-asset data, leveraging gRPC for inter-service communication and Merkle proofs to ensure the integrity of historical price feeds. This system allows for trustless verification of price data, enabling accurate and reliable data access for decentralized applications (dApps), exchanges, and other financial institutions. It aims to address the common problem of unreliable and potentially manipulated price data sources by creating a verifiable and transparent record.

The project consists of several microservices that work together to collect, aggregate, and secure price data. Data is sourced from various cryptocurrency exchanges and aggregated using a weighted average algorithm. The aggregated data is then anchored to a Merkle tree, creating a cryptographically secure history of price feeds. gRPC is used for communication between the microservices, ensuring efficient and reliable data transfer. This architecture allows for horizontal scalability and fault tolerance, critical for maintaining a robust and resilient data infrastructure.

ModernQuantumVault offers a significant advantage over traditional centralized data providers by providing verifiable data integrity. The use of Merkle proofs allows users to independently verify the historical accuracy of price data, eliminating the need to trust a single source. This transparency fosters greater trust and confidence in the data, making it suitable for use in critical applications where data integrity is paramount. Furthermore, the decentralized nature of the system mitigates the risk of data manipulation or single points of failure.

## Key Features

*   **gRPC-Based Microservices Architecture:** Leverages gRPC for high-performance, type-safe communication between data ingestion, aggregation, and storage services. This ensures efficient data transfer and a well-defined API.

*   **Decentralized Data Aggregation:** Aggregates data from multiple cryptocurrency exchanges using a weighted average algorithm. The weighting can be configured to prioritize more reputable exchanges.

*   **Merkle Tree Anchoring:** Anchors historical price data to a Merkle tree, providing a tamper-evident record of price feeds. Each block of data contains the root hash of the Merkle tree at that point in time.

*   **Verifiable Price Feeds:** Provides Merkle proofs that allow users to independently verify the integrity of historical price data. This enables trustless verification of price accuracy. The proofs can be verified against the published Merkle root.

*   **Configurable Data Sources:** Allows for easy addition and removal of data sources (cryptocurrency exchanges). This flexibility ensures that the system can adapt to changes in the market landscape. The configuration is managed through environment variables.

*   **Fault Tolerance:** Designed with a microservices architecture to ensure fault tolerance. Individual service failures do not impact the overall system functionality.

*   **REST API:** Exposes a REST API layer built on top of the gRPC backend for easy querying of current and historical price data, as well as Merkle proofs. This enables integration with a wide range of applications.

## Technology Stack

*   **Go:** The primary programming language for all microservices, chosen for its performance, concurrency support, and strong typing.
*   **gRPC:** A high-performance, open-source universal RPC framework for inter-service communication, providing efficient and reliable data transfer.
*   **Protocol Buffers:** Used to define the gRPC service interfaces and message formats, ensuring type safety and efficient serialization.
*   **Merkle Trees:** Used to create a tamper-evident record of historical price data, enabling verifiable data integrity. The `merkletree` library in Go is used for this implementation.
*   **PostgreSQL:** A robust and scalable relational database used to store aggregated price data and Merkle tree hashes.
*   **REST API (using Gin/Echo):** A RESTful API layer for external querying of price data and Merkle proofs.

## Installation

1.  **Prerequisites:**
    *   Go (version 1.18 or higher)
    *   PostgreSQL (version 12 or higher)
    *   Docker (optional, for containerization)

2.  **Clone the repository:**
    git clone https://github.com/jjfhwang/ModernQuantumVault.git
    cd ModernQuantumVault

3.  **Set up the PostgreSQL database:**
    *   Create a database named `modernquantumvault`.
    *   Create a user with the necessary permissions to access the database.

4.  **Configure environment variables (see Configuration section below).**

5.  **Build the project:**
    go build ./...

6.  **Run the microservices:**
    *   Each microservice (data ingestion, aggregation, storage, API) should be run separately.
    *   Example: ./data-ingestion

## Configuration

The following environment variables are required for each microservice:

*   `DATABASE_HOST`: The hostname or IP address of the PostgreSQL server.
*   `DATABASE_PORT`: The port number of the PostgreSQL server.
*   `DATABASE_USER`: The username for the PostgreSQL database.
*   `DATABASE_PASSWORD`: The password for the PostgreSQL database.
*   `DATABASE_NAME`: The name of the PostgreSQL database.
*   `GRPC_PORT`: The port number for the gRPC server.
*   `REST_PORT`: The port number for the REST API server (for the API service only).
*   `EXCHANGE_API_KEYS`: (For Data Ingestion Service only) JSON containing the API keys for the exchanges being used (e.g., `{"exchange1": "API_KEY_EXCHANGE_1", "exchange2": "API_KEY_EXCHANGE_2"}`).
*   `AGGREGATION_WEIGHTS`: (For Aggregation Service only) JSON containing weights for each exchange (e.g., `{"exchange1": 0.6, "exchange2": 0.4}`).

These environment variables can be set directly in your shell or using a `.env` file.

## Usage

The primary way to interact with ModernQuantumVault is through the REST API. The API provides endpoints for:

*   **Getting the current price of a crypto-asset:**
    `GET /price/{symbol}`
    Returns the current aggregated price for the specified symbol (e.g., BTCUSD).

*   **Getting historical price data:**
    `GET /historical/{symbol}?timestamp={timestamp}`
    Returns the historical price for the specified symbol at the given timestamp (Unix timestamp).

*   **Getting a Merkle proof for historical price data:**
    `GET /proof/{symbol}?timestamp={timestamp}`
    Returns the Merkle proof for the specified symbol at the given timestamp. The returned JSON includes the Merkle root, the leaf data, and the proof path.

Example using `curl`:

Retrieve the current price of BTCUSD:
curl http://localhost:{REST_PORT}/price/BTCUSD

Retrieve the Merkle proof for BTCUSD at timestamp 1678886400:
curl http://localhost:{REST_PORT}/proof/BTCUSD?timestamp=1678886400

## Contributing

We welcome contributions to ModernQuantumVault! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with thorough documentation.
4.  Submit a pull request with a detailed description of your changes.
5.  Ensure all tests pass before submitting your pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/ModernQuantumVault/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to acknowledge the open-source community for their contributions to the Go ecosystem and the libraries we have used in this project. Specifically, we thank the developers of the gRPC, Protocol Buffers, and PostgreSQL drivers for Go.