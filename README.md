<a name="readme-top"></a>

<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->



<!-- PROJECT LOGO -->

<br />
<div align="center">
  <picture width="160px" align="center">
      <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/terramate-io/brand/5a799813d429116741243b9b06a9f034a3991bf3/darkmode/stamp.svg">
      <img alt="Terramate" src="https://raw.githubusercontent.com/terramate-io/brand/5a799813d429116741243b9b06a9f034a3991bf3/whitemode/stamp.svg" width="160px" align="center">
    </picture>

<h3 align="center">How to Use Atlantis and Terramate Cloud on a Local Linux Machine</h3>

<p align="center">
    An effective local setup for managing infrastructure as code using Atlantis and Terramate Cloud.
    <br />
    <br />
    <a href="https://github.com/RocketRene/terramate-atlantis"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/RocketRene/terramate-atlantis/pull/8">View Demo</a>
    ·
    <a href="https://github.com/RocketRene/terramate-atlantis/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/RocketRene/terramate-atlantis/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project


![Atlantis PR Automation](<images/Pasted image.png>)


This project provides a seamless integration setup for Atlantis and Terramate Cloud, designed to run on local Linux environments. By utilizing Atlantis, a tool that makes it easier to use Terraform in a team setting, along with Terramate Cloud, which simplifies Terraform state management, this setup aims to streamline the development and deployment of infrastructure as code. The integration allows developers to test locally before deploying to the cloud.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [Terramate Cloud](https://terramate.io)
* [Atlantis](https://runatlantis.io)
* [GitHub](https://github.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To get this integration up and running locally, follow these simple steps.

### Prerequisites

Ensure you have the following tools installed and configured on your system:

- Terraform
- Atlantis `brew install atlantis`
- AWS-Vault `brew install aws-vault`
- ngrok 
  - `wget https://bin.equinox.io/c/bNyj1mQVY4c/ngrok-v3-stable-linux-amd64.tgz`
  - `sudo tar xvzf ./ngrok-v3-stable-linux-amd64.tgz -C /usr/local/bin`

- Terramate CLI `brew install terramate`
- GitHub CLI `brew install gh` 
- Go Compiler `brew install go`

### Installation


1. **Clone the repository:**
   ```sh
   git clone https://github.com/RocketRene/terramate-atlantis.git 
   ```
2. **Navigate to the atlanis  directory in the repo :**
   ```sh
    cd terramate-atlantis/atlantis
    ```
3. **Set up the environment variables in the `.env` file:**
   ```sh
   cp example.env .env
   nano .env
   ```
4. **Authenticate with AWS-Vault:**
   ```sh
   aws-vault exec <profile-name> 
   ```  
5. **Authenticate with Terramate CLI:**
    ```sh
    terramate cloud login
    ```
6. **Create a Webkook in the repo settings:**
   
    You have to manually create a webhook in the repo settings, use `example.com` as the URL and `application/json` as the content type. The webhook should send everything.
    Use the secret from the `.env` file as the secret.
 
    
7. **Install the dependencies for the custom setup Tool:**
    ```sh
    go mod tidy
    ```
8. **Start the Atlantis server:**
    ```sh
    go run setupAtlantis.go
    ```
    The setupAtlantis.go script automates the setup of the Atlantis server integrated with GitHub, utilizing ngrok to expose Atlantis to the web. It starts by launching an ngrok tunnel and retrieving the public URL, configures a GitHub client with authentication, updates GitHub webhooks to point to the new ngrok URL, and finally starts the Atlantis server using environmental variables for configuration. 



<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

To utilize the automated Atlantis setup and test its functionality, follow these steps from the root of your repository:

1. Ensure you are in the repository root:
If you are in the 'atlantis' directory, navigate back to the repository root:

```sh
cd ..
```

2. Run the test script:
Execute the testAtlantis.go script to automatically handle various Git operations and create a pull request:

```sh
    go run testAtlantis.go
```

This script (`testAtlantis.go`) performs several automated tasks including switching to the main branch, updating it, generating a unique branch name, and making a minimal change to your cloud infrastructure by randomly altering an S3 bucket name in the main.tf file. After modifying the configuration, the script commits the changes and pushes them back to GitHub. It concludes by creating a new pull request using the GitHub CLI, which is crucial for testing and demonstrating the dynamic configuration capabilities of the Atlantis setup in a real-world scenario.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [ ] Create a custom `atlantis.yaml` generator
- [ ] Build a Docker image that integrates both Terramate and Atlantis for simplified deployment


See the [open issues](https://github.com/RocketRene/terramate-atlantis/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the Apache License Version 2.0. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

[René Kuhn](mailto:rene.kuhn@terramate.io)

<p align="right">(<a href="#readme-top">back to top</a>)</p>




