# What's this?
Little tool written in Go to merge several individual ssh configuration files to a single configuration. 
Bash equivalent: 
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