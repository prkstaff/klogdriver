# klogdriver
Parse kubernetes `kubectl logs` stackdriver json format for better local visualization.

## Installing
### 1 - Download the latest release
```
sudo mv klogdriver /usr/local/bin
```

### 2 - Create and alias
```
function klog () { kubectl logs -f --timestamps $@ | klogdriver ; }
```
