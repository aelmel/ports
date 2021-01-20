print('Start #################################################################');

db.createUser({
  user: 'port-user',
  pwd: 'port-password',
  roles: [
    {
      role: 'readWrite',
      db: 'ports',
    },
  ],
});

db.createCollection('details');
