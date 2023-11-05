# phpipamsync

Small script synchronizes phpipam hosts with pi-hole resolvers. 

## Use case

Phpipam considered source of truth for DNS resolution. Network admins create hostnames in phpipam. 
Pihole used as a DNS-server. 
To synchronize phpipam hostanmes with pihole DNS-resolver 
```
phpipamsync get pi-hole
```

## phpipamsync config

Location ~/.phpipamsync/config

```
ipam_site_url: https://192.168.169.170:8670/                    # IP and port where phpipam exposes API endpoints
ipam_app_id: phpipamsync                                        # phpipam app id
ipam_app_code: AVyzJl7ceWJ5WrbKZ1QjAIgNdMSggw12                 # phpipam app code
ipam_subnet: 192.168.168.0/24                                   # ipam subnet to synchronize

domain: home                                                    # domain to add to hostnames. For example (home will add .home to all hostnames)
pi_hole: /home/ikuchin/docker/pi.hole/etc-pihole/custom.list    # full path to pi-hole custom.list file 
```

**ATTENTION** pi-hole custom file will be rewritten, not appended.
