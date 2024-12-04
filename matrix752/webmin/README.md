# Webmin 相關路徑

## 1. Webmin 根目錄
```
/usr/libexec/webmin
```

## 2. Webmin 設定檔
```
/etc/webmin/miniserv.conf
```

## 3. Webmin 啟動與停用
```
/etc/webmin/start
/etc/webmin/stop
```

## 4. Check Webmin PID
```
cat /var/webmin/miniserv.pid
```

## 5. 系統開機自動啟動 Webmin
啟動腳本的軟連結為 `/etc/rc5.d/S99webmin`

真正的啟動腳本則是 `/etc/init.d/webmin`

啟動腳本會讀取 `/etc/webmin/miniserv.conf` 設定檔

然後執行 `/usr/libexec/webmin/start` 啟動 miniserv

miniserv 執行時的 PID 會寫入 `/var/webmin/miniserv.pid`

