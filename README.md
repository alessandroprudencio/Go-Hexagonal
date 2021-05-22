Crud de produtos utilizando a arquitetura hexagonal com go lang, com opção de acesso via cli como também via servidor web

### Via CLI 

Flags:

``-a, --action string    Enable /  Disable a product (default "enable")``

``-h, --help             help for cli``

``-i, --id string        Product ID``

``-p, --price float      Product price``

``-n, --product string   Product name``
    
Cadastrar Produto

    go run main.go cli -a=create -n=Product cli -p=15.00

Buscar Produto

    go run main.go cli -a=get -i=edc49fdd-f5f3-4dca-a4a4-8ffcfd744740

### Via Web Server :

Buscar Produto

    localhost:8080/product/452a8fd9-9e0b-400e-815e-b03a708d19c6
