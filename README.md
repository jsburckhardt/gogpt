# gogpt

GoGPT is a cli tool to work with OpenAI / Azure OpenAI models.

## Installation

- If you know the specific version you would like to install

  ```bash
  VERSION=v1.0.0
  wget -q -O - https://github.com/jsburckhardt/gogpt/releases/download/$VERSION/install.sh | bash
  ```

- If you want to pull the latest. You'll need jq

  ```bash
  VERSION=$(curl -sL https://api.github.com/repos/jsburckhardt/gogpt/releases/latest | jq -r ".tag_name")
  wget -q -O - https://github.com/jsburckhardt/gogpt/releases/download/$VERSION/install.sh | bash
  ```

## Config

you can create a gogpt.conf or .env or export variables

```.env
OPENAI_API_TYPE="azure"
OPENAI_API_MODEL="<model>"
OPENAI_API_KEY="<your key>"
OPENAI_API_HOST="https://<lab>.openai.azure.com/"
AZURE_OPENAI_ENGINE="<engine>"
```

```bash
export OPENAI_API_TYPE="azure"
export OPENAI_API_MODEL="<model>"
export OPENAI_API_KEY="<your key>"
export OPENAI_API_HOST="https://<lab>.openai.azure.com/"
export AZURE_OPENAI_ENGINE="<engine>"
```

or for openai

```.env
OPENAI_API_KEY="<your key>"
```

```bash
export OPENAI_API_KEY="<your key>"
```

## Inspired

- [sgpt](https://github.com/TheR1D/shell_gpt)
