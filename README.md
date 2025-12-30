Projet : Calculatrice Web en Go
Une application web simple et robuste permettant d'effectuer des opérations arithmétiques de base, développée avec le langage Go et les templates HTML standards.

🚀 Fonctionnalités
-**Opérations supportées** : Addition, Soustraction, Multiplication et Division.

-**Interface intuitive** : Saisie via formulaire numérique et sélection d'opération par boutons radio.

-**Gestion d'erreurs robuste** :

-**Validation des entrées** (nombres valides uniquement).

-**Détection et blocage** de la division par zéro.

-**Vérification** côté client (HTML5) et côté serveur (Go).

🛠️ Technologies utilisées
Backend : Go (Golang)

Frontend : HTML / CSS (Design moderne et responsive)

Architecture : Pattern MVC simplifié (Model/Structure, View/Template, Controller).

📁 Structure du projet
proj2/
├── controller/
│   └── controller.go   # Logique métier et validation
├── router/
│   └── router.go       # Configuration des routes et fichiers statiques
├── structure/
│   └── structure.go    # Définition de la structure PageData
├── static/
│   └── style.css       # Design et mise en page
├── template/
│   └── home.html       # Interface utilisateur (Go Templates)
└── main.go             # Point d'entrée de l'application
⚙️ Installation et Lancement
Cloner le projet :

Bash

git clone <url-du-depot>
cd proj2
Lancer le serveur :

Bash

go run main.go
Accéder à l'application : Ouvrez votre navigateur et accédez à l'adresse suivante : http://localhost:8080/home
