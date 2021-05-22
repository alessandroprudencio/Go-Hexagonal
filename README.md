Crud de produtos utilizando a arquitetura hexagonal com go lang

Comandos via CLI 

Flags:

``-a, --action string    Enable /  Disable a product (default "enable")``

``-h, --help             help for cli``

``-i, --id string        Product ID``

``-p, --price float      Product price``

``-n, --product string   Product name``
    
Uso:

Cadastrar Produto

    go run main.go cli -a=create -n=Product cli -p=15.00

Buscar Produto

    go run main.go cli -a=get -i=edc49fdd-f5f3-4dca-a4a4-8ffcfd744740
