## You need to stop any service running 80

    sudo netstat -tulpn | grep 80



    sudo ss -tulpn | grep 80


    sudo fuser 80/tcp

Try killing the process  IDs

If None of the these above work please  remove the apache2 service

    sudo apt remove apache2