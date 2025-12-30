# 🧮 Calculatrice Web en Go

Une application web robuste et ergonomique permettant d'effectuer des opérations arithmétiques, développée en **Go** avec une architecture modulaire.

---

## 🚀 Fonctionnalités
* **Opérations arithmétiques** : Addition, Soustraction, Multiplication et Division.
* **Interface intuitive** : Saisie numérique sécurisée et sélection d'opérations via boutons radio.
* **Gestion d'erreurs avancée** :
    * Validation des types de données (numérique uniquement).
    * Blocage de la **division par zéro** avec message d'alerte.
    * Double validation : Côté client (HTML) et côté serveur (Go).

## 🛠️ Technologies utilisées
* **Backend** : Go (Golang) - Utilisation de `net/http` et `html/template`.
* **Frontend** : HTML / CSS (Design moderne sans framework).
* **Architecture** : Pattern MVC simplifié (Model / View / Controller).

---

## 📁 Structure du projet
```text
calculatrice/
├── controller/
│   └── controller.go   # Logique métier et gestion des requêtes
├── router/
│   └── router.go       # Configuration des routes et ressources statiques
├── structure/
│   └── structure.go    # Définition du modèle de données (PageData)
├── static/
│   └── style.css       # Mise en page et design responsive
├── template/
│   └── home.html       # Interface utilisateur via Go Templates
└── main.go             # Point d'entrée de l'application

🛠️ Installation et Lancement
1. Cloner le dépôt :
    git clone https://github.com/joffretkevin727/calculatrice.git

2. Lancer le serveur :
    cd .\calculatrice\
    go run main.go

3. Accéder à l'application :
    Rendez-vous sur http://localhost:8080/home
