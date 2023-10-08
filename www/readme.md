#------------------------------------------

ln -s /etc/nginx/sites-available/umlaphi.ricsa.co.za.conf /etc/nginx/sites-enabled/
sudo nginx -t && sudo nginx -s reload
cd /opt/letsencrypt

./certbot-auto --config /etc/letsencrypt/configs/pmis2.easipath.com.conf certonly

certbot certonly --nginx

        access_log /var/log/nginx/access.log;
        error_log /var/log/nginx/error.log;

return 301 https://cclms2.easipath.com$request_uri;



https://www.youtube.com/watch?v=MyVOLRT7PVs&t=16s