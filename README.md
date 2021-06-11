# What's this?
Little tool written in Go to merge several individual ssh configuration files to a single configuration.
```sh
cat ~/.ssh/config_* >> ~/.ssh/config
```

## Usage
```go
ssh-merge-config ~/.ssh/config.d/ ~/.ssh/config
```

## Output
```sh
~  ssh-merge-config ~/.ssh/config.d/ ~/.ssh/config
done
```

## Shell equivalent
You can achieve the same by doing the following in the console:
```sh
cat ~/.ssh/config_* >> ~/.ssh/config
```