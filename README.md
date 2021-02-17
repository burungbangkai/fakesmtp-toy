# FAKE SMTP
As the name describe, it is a fake SMTP server that trap any email send to it.
It consist of 2 service:

1. SMTP Multiplexer

    A service that handle SMTP service. it multiplex 1 SMTP to multiple SMTP account. in other words an SMTP server that handle multiple inboxes

2. Inboxes Management

    A service that manage inboxes; such as create, update delete. an inbox will have a name, user, and password

