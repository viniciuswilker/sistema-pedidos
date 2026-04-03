CREATE DATABASE IF NOT EXISTS sistemapedidosdb;
USE sistemapedidosdb;

DROP TABLE IF EXISTS cliente;
CREATE TABLE cliente (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) UNIQUE, 
    email VARCHAR(100) UNIQUE,
    senha VARCHAR(255),
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

DROP TABLE IF EXISTS categorias;
CREATE TABLE categorias (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL
) ENGINE=InnoDB;

DROP TABLE IF EXISTS produtos;
CREATE TABLE produtos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    categoria_id INT NOT NULL,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    preco DECIMAL(10, 2) NOT NULL,
    disponivel BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS pedidos;
CREATE TABLE pedidos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_id INT,
    data_pedido TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status ENUM('pendente', 'em_preparo', 'pronto', 'entregue', 'cancelado') DEFAULT 'pendente',
    total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (cliente_id) REFERENCES cliente(id)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS itens_pedido;
CREATE TABLE itens_pedido (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pedido_id INT NOT NULL,
    produto_id INT NOT NULL,
    quantidade INT NOT NULL,
    preco_unitario DECIMAL(10, 2) NOT NULL, 
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id) ON DELETE CASCADE,
    FOREIGN KEY (produto_id) REFERENCES produtos(id)
) ENGINE=InnoDB;

CREATE INDEX idx_status_pedido ON pedidos(status);
CREATE INDEX idx_produto_disponivel ON produtos(disponivel);


DROP TABLE IF EXISTS funcionarios;
CREATE TABLE funcionarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    usuario VARCHAR(30) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha VARCHAR(255) NOT NULL,
    cargo ENUM('admin', 'operador') DEFAULT 'operador',
    ativo BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

ALTER TABLE pedidos ADD COLUMN atualizado_por INT;
ALTER TABLE pedidos ADD FOREIGN KEY (atualizado_por) REFERENCES funcionarios(id);


DROP TABLE IF EXISTS notas;

CREATE TABLE notas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pedido_id INT UNIQUE NOT NULL, 
    numero_nota VARCHAR(50) UNIQUE NOT NULL, 
    data_emissao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valor_total_nota DECIMAL(10, 2) NOT NULL,
    dados_snapshot JSON NOT NULL,
    
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id) ON DELETE RESTRICT
) ENGINE=InnoDB;

CREATE INDEX idx_numero_nota ON notas(numero_nota);







-- CATEGORIAS
insert into categorias (nome) values ("hamburguer");
insert into categorias (nome) values ("bebidas");
insert into categorias (nome) values ("doces");


-- PRODUTOS
insert into produtos (categoria_id, nome, descricao, preco, disponivel) values (2, "Coca Cola 350ml" , "Lata de Coca Cola 350ml", 4.00, true);

insert into produtos (categoria_id, nome, descricao, preco, disponivel) values (2, "Franta 350ml" , "Franta 350ml", 3.99, true);

insert into produtos (categoria_id, nome, descricao, preco, disponivel) values (1, "X TUDO" , "X tudo com queijo, hamburguer duplo, bacon, etc.", 25.00, true);
