# Proof of Concept – Event driven architecture with Kafka

## Description du projet

Simulation d'un système de e-commerce basé sur une architecture pilotée par les événements utilisant Apache Kafka pour la communication entre les services.

## Liste des tâches

- [ ] La publication d'événements par deux ou plusieurs composants producteurs
- [x] Producteur 1 : Service de gestion des commandes (Orders)
- [ ] Producteur 2 : A décider
- [ ] La consommation et le traitement de ces événements par deux ou plusieurs composants
- [x] Consommateur 1 : Service de d'iventaire et de notification
- [ ] Consommateur 2 : A décider
      consommateurs
- [ ] Des réactions observables suite au traitement des événements (par exemple : mise à jour
      d'état, persistance en base de données, notification, journalisation, etc.)
  - [ ] Docker avec une base de données de l'inventaire (PostgreSQL, MongoDB, etc.) pour la persistance des données
  - [ ] Journalisation des événements traités dans un fichier ou une base de données
- [ ] Un déclenchement manuel clair via interface utilisateur
- [ ] Interface utilisateur simple web pour initier des actions dans le système

- [ ] Fournir un diagramme d'architecture clair illustrant :
  - [ ] Les services/processus impliqués
  - [ ] Le(s) topic(s) Kafka utilisé(s)
  - [ ] Les flux de publication et de consommation
  - [ ] La séquence des événements du stimulus initial à la réaction finale
  - [ ] Description conscise expliquement le fonctionnement global du système

- [ ] Document d'architecture et instructions
  - [ ] Partie 1 Architecture voir le doc sur moodle
  - [ ] Partie 2 Instructions de démarrage

- [ ] Réalisation d'une vidéo de démonstration

## Extrait vidéo

[![Démonstration vidéo]()]()

## Objectif

TODO:

---

## Tactiques mises en œuvre

TODO

---

## Design de l’application

TODO

### Composants

| Composant | Description | Port | Rôle |
| --------- | ----------- | ---- | ---- |

TODO

### Schéma d’architecture

![img](documentation/architecture.png)

---

## Topics

### Topic 1

### Topic 2

---

## Exécution du projet en local

Il est obligatoire d'avoir docker sur sa machine.

Pour lancer l'ensemble du projet (a exécuter à la racine du projet):

```bash
docker compose up --build
```

Cette commande va:

- TODO
