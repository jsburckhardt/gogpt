# gogpt

GoGPT is a cli tool to work with OpenAI / Azure OpenAI models.

## Config

you can create a gogpt.conf or .env or export variables

```txt
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

```txt
OPENAI_API_KEY="<your key>"
```

## Inspired

- [sgpt](https://github.com/TheR1D/shell_gpt)
