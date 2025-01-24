\c wolfgres_db

-- Schema wfg
CREATE SCHEMA wfg;


-- Table Customer
CREATE TABLE wfg.customer(
 customer_id BIGSERIAL PRIMARY KEY,
 name VARCHAR(255),
 address VARCHAR(255),
 website VARCHAR(255),
 credit_limit double precision
);

-- Table Operation
CREATE TABLE wfg.operation(
 operation_id BIGSERIAL PRIMARY KEY,
 name VARCHAR(50),
 description TEXT
);

-- Table account_type
CREATE TABLE wfg.account_type(
 account_type_id BIGSERIAL PRIMARY KEY,
 name VARCHAR(50),
 description TEXT
);

-- Table Account
CREATE TABLE wfg.account(
 account_id BIGSERIAL PRIMARY KEY,
 customer_id BIGINT,
 account_type_id BIGINT,
 balace DOUBLE PRECISION,
 CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES wfg.customer(customer_id),
 CONSTRAINT fk_account_type FOREIGN KEY(account_type_id) REFERENCES wfg.account_type(account_type_id)
);

-- Table Transaction
CREATE TABLE wfg.transaction(
 transaction_id BIGSERIAL PRIMARY KEY,
 account_id BIGINT,
 operation_id BIGINT,
 mount DOUBLE PRECISION,
 date TIMESTAMP,
 CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES wfg.account(account_id),
 CONSTRAINT fk_operation FOREIGN KEY(operation_id) REFERENCES wfg.operation(operation_id)
);

-- INSERTS Customer
INSERT INTO wfg.customer VALUES (1, 'Tazz', 'Maple Wood Park 28', 'tazz.global', 75161.4282);
INSERT INTO wfg.customer VALUES (2, 'Topdrive', 'Kenwood Trail 13', 'topdrive.info', 91480.2339);
INSERT INTO wfg.customer VALUES (3, 'Yakijo', 'Cottonwood Street 10', 'yakijo.biz', 625851.1481);
INSERT INTO wfg.customer VALUES (4, 'Browsetype', 'Browning Parkway 83', 'browsetype.site', 914489.6477);
INSERT INTO wfg.customer VALUES (5, 'Gabtype', 'Mosinee Drive 26', 'gabtype.support', 207078.4459);
INSERT INTO wfg.customer VALUES (6, 'Photobug', 'Aberg Pass 93', 'photobug.solutions', 383007.8109);
INSERT INTO wfg.customer VALUES (7, 'Rhynoodle', 'Dunning Court 91', 'rhynoodle.support', 657781.417);
INSERT INTO wfg.customer VALUES (8, 'Skidoo', 'Park Meadow Drive 3', 'skidoo.org', 302429.3872);
INSERT INTO wfg.customer VALUES (9, 'Zoomdog', 'Northfield Drive 99', 'zoomdog.biz', 117979.0628);
INSERT INTO wfg.customer VALUES (10, 'Linklinks', 'Declaration Crossing 62', 'linklinks.biz', 245552.1986);
INSERT INTO wfg.customer VALUES (11, 'Gabtune', 'Sunfield Street 99', 'gabtune.app', 809597.2417);
INSERT INTO wfg.customer VALUES (12, 'Leenti', 'Sherman Park 20', 'leenti.global', 304653.429);
INSERT INTO wfg.customer VALUES (13, 'Katz', 'Maryland Junction 42', 'katz.global', 243237.3976);
INSERT INTO wfg.customer VALUES (14, 'Mycat', 'Saint Paul Lane 81', 'mycat.xyz', 36859.9406);
INSERT INTO wfg.customer VALUES (15, 'Bluezoom', 'Sutherland Trail 12', 'bluezoom.global', 833904.4564);
INSERT INTO wfg.customer VALUES (16, 'Topdrive', 'Fuller Circle 92', 'topdrive.site', 771903.7015);
INSERT INTO wfg.customer VALUES (17, 'Ooba', 'Pearson Circle 88', 'ooba.me', 260631.8087);
INSERT INTO wfg.customer VALUES (18, 'Jayo', 'Kennedy Court 10', 'jayo.global', 2894.6094);
INSERT INTO wfg.customer VALUES (19, 'Vidoo', 'Crownhardt Crossing 80', 'vidoo.xyz', 525691.0929);
INSERT INTO wfg.customer VALUES (20, 'Izio', 'Dixon Court 73', 'izio.online', 47191.9153);
INSERT INTO wfg.customer VALUES (21, 'Meejo', 'Texas Lane 26', 'meejo.net', 706872.4062);
INSERT INTO wfg.customer VALUES (22, 'Camimbo', 'Raven Junction 29', 'camimbo.solutions', 361908.5547);
INSERT INTO wfg.customer VALUES (23, 'LiveZ', 'Holy Cross Court 97', 'livez.io', 770270.1435);
INSERT INTO wfg.customer VALUES (24, 'Gigabox', 'Doe Crossing Circle 88', 'gigabox.xyz', 564377.287);
INSERT INTO wfg.customer VALUES (25, 'Bluejam', 'Larry Lane 43', 'bluejam.co', 565235.2394);
INSERT INTO wfg.customer VALUES (26, 'Twitternation', 'Haas Drive 3', 'twitternation.me', 755682.4315);
INSERT INTO wfg.customer VALUES (27, 'Skippad', 'Myrtle Terrace 46', 'skippad.site', 66123.4238);
INSERT INTO wfg.customer VALUES (28, 'Shufflebeat', 'Derek Lane 2', 'shufflebeat.me', 160089.7771);
INSERT INTO wfg.customer VALUES (29, 'Oyonder', 'Mandrake Hill 11', 'oyonder.biz', 224872.7999);
INSERT INTO wfg.customer VALUES (30, 'Wikivu', 'Sachs Point 81', 'wikivu.support', 57878.0629);
INSERT INTO wfg.customer VALUES (31, 'Mymm', 'Huxley Circle 45', 'mymm.xyz', 481484.9678);
INSERT INTO wfg.customer VALUES (32, 'Vinder', 'Farwell Court 35', 'vinder.support', 492894.1295);
INSERT INTO wfg.customer VALUES (33, 'Kare', 'Graedel Court 10', 'kare.online', 943438.1552);
INSERT INTO wfg.customer VALUES (34, 'Tanoodle', 'Springs Junction 32', 'tanoodle.support', 846859.2877);
INSERT INTO wfg.customer VALUES (35, 'Fanoodle', 'Kinsman Circle 73', 'fanoodle.net', 352332.925);
INSERT INTO wfg.customer VALUES (36, 'Linkbuzz', 'Spohn Road 10', 'linkbuzz.app', 599998.9219);
INSERT INTO wfg.customer VALUES (37, 'Gabtype', 'Spaight Street 42', 'gabtype.org', 617633.5585);
INSERT INTO wfg.customer VALUES (38, 'Fliptune', 'Waywood Lane 36', 'fliptune.support', 383508.1552);
INSERT INTO wfg.customer VALUES (39, 'Abata', 'Hovde Junction 13', 'abata.biz', 932276.0311);
INSERT INTO wfg.customer VALUES (40, 'Jabberstorm', 'Menomonie Trail 13', 'jabberstorm.support', 725034.1797);
INSERT INTO wfg.customer VALUES (41, 'Roombo', 'Kenwood Pass 22', 'roombo.solutions', 294882.391);
INSERT INTO wfg.customer VALUES (42, 'Eamia', 'Stone Corner Point 62', 'eamia.site', 898699.9451);
INSERT INTO wfg.customer VALUES (43, 'Rhyloo', 'East Road 93', 'rhyloo.site', 64415.6816);
INSERT INTO wfg.customer VALUES (44, 'Flipstorm', 'Armistice Lane 92', 'flipstorm.me', 193057.3915);
INSERT INTO wfg.customer VALUES (45, 'Flipopia', 'Karstens Center 97', 'flipopia.io', 808282.0938);
INSERT INTO wfg.customer VALUES (46, 'Jabbersphere', 'Tomscot Way 72', 'jabbersphere.app', 958346.1544);
INSERT INTO wfg.customer VALUES (47, 'Yombu', 'Russell Place 43', 'yombu.com', 571353.6642);
INSERT INTO wfg.customer VALUES (48, 'Realcube', 'Welch Lane 22', 'realcube.org', 878349.761);
INSERT INTO wfg.customer VALUES (49, 'Lazzy', 'Daystar Pass 75', 'lazzy.io', 534216.528);
INSERT INTO wfg.customer VALUES (50, 'Feedbug', 'Lawn Parkway 16', 'feedbug.co', 493584.9331);
INSERT INTO wfg.customer VALUES (51, 'Bluezoom', 'Melrose Junction 8', 'bluezoom.support', 509999.8037);
INSERT INTO wfg.customer VALUES (52, 'Camido', 'Northwestern Lane 78', 'camido.com', 692874.3964);
INSERT INTO wfg.customer VALUES (53, 'Voolith', 'New Castle Plaza 77', 'voolith.com', 201011.3951);
INSERT INTO wfg.customer VALUES (54, 'Topicblab', 'Prentice Road 85', 'topicblab.info', 200989.5316);
INSERT INTO wfg.customer VALUES (55, 'Skajo', 'Clarendon Center 96', 'skajo.net', 784693.9371);
INSERT INTO wfg.customer VALUES (56, 'Wordtune', 'Northfield Alley 51', 'wordtune.online', 451517.742);
INSERT INTO wfg.customer VALUES (57, 'Voomm', 'Porter Terrace 67', 'voomm.org', 153457.7273);
INSERT INTO wfg.customer VALUES (58, 'Dablist', 'Towne Road 42', 'dablist.net', 648976.8488);
INSERT INTO wfg.customer VALUES (59, 'Einti', 'Colorado Plaza 4', 'einti.io', 996370.1705);
INSERT INTO wfg.customer VALUES (60, 'Pixonyx', 'Comanche Circle 70', 'pixonyx.org', 422876.7421);
INSERT INTO wfg.customer VALUES (61, 'Oyoloo', 'Jenna Point 0', 'oyoloo.io', 292409.1287);
INSERT INTO wfg.customer VALUES (62, 'Skaboo', 'Hovde Center 95', 'skaboo.online', 623686.2538);
INSERT INTO wfg.customer VALUES (63, 'Eimbee', 'Fordem Street 73', 'eimbee.co', 754456.734);
INSERT INTO wfg.customer VALUES (64, 'Twimbo', 'Boyd Pass 10', 'twimbo.xyz', 926809.5769);
INSERT INTO wfg.customer VALUES (65, 'Kaymbo', '2nd Junction 86', 'kaymbo.online', 262761.8753);
INSERT INTO wfg.customer VALUES (66, 'Devshare', 'Weeping Birch Drive 94', 'devshare.me', 758960.9369);
INSERT INTO wfg.customer VALUES (67, 'Realmix', 'Meadow Ridge Way 84', 'realmix.xyz', 155572.2785);
INSERT INTO wfg.customer VALUES (68, 'Livetube', 'Hallows Pass 56', 'livetube.site', 673071.7883);
INSERT INTO wfg.customer VALUES (69, 'Avamm', 'Magdeline Court 70', 'avamm.global', 486097.8566);
INSERT INTO wfg.customer VALUES (70, 'Flipbug', 'Sommers Lane 21', 'flipbug.net', 166597.9953);
INSERT INTO wfg.customer VALUES (71, 'Oyoba', 'Duke Crossing 45', 'oyoba.biz', 177912.0145);
INSERT INTO wfg.customer VALUES (72, 'Zoombox', 'Spenser Alley 14', 'zoombox.io', 805702.0206);
INSERT INTO wfg.customer VALUES (73, 'Ntag', 'Dayton Terrace 30', 'ntag.io', 748622.9993);
INSERT INTO wfg.customer VALUES (74, 'Twiyo', 'Cottonwood Pass 96', 'twiyo.app', 110163.6391);
INSERT INTO wfg.customer VALUES (75, 'Photolist', 'Portage Parkway 74', 'photolist.io', 137226.0315);
INSERT INTO wfg.customer VALUES (76, 'Ooba', 'Carey Pass 25', 'ooba.online', 748390.267);
INSERT INTO wfg.customer VALUES (77, 'Ailane', 'Doe Crossing Plaza 74', 'ailane.com', 671388.0817);
INSERT INTO wfg.customer VALUES (78, 'Eazzy', 'Grayhawk Trail 70', 'eazzy.online', 62858.4158);
INSERT INTO wfg.customer VALUES (79, 'Realmix', 'Kenwood Crossing 38', 'realmix.xyz', 662217.6376);
INSERT INTO wfg.customer VALUES (80, 'JumpXS', 'Forest Pass 76', 'jumpxs.solutions', 35635.2167);
INSERT INTO wfg.customer VALUES (81, 'Linkbuzz', 'Arrowood Avenue 24', 'linkbuzz.com', 815700.8603);
INSERT INTO wfg.customer VALUES (82, 'Tagchat', 'Southridge Lane 20', 'tagchat.info', 713457.0183);
INSERT INTO wfg.customer VALUES (83, 'Cogilith', 'Corry Junction 29', 'cogilith.org', 899857.0416);
INSERT INTO wfg.customer VALUES (84, 'Kwilith', 'Browning Way 29', 'kwilith.me', 965574.4198);
INSERT INTO wfg.customer VALUES (85, 'Riffpedia', 'Roxbury Terrace 53', 'riffpedia.online', 326292.0283);
INSERT INTO wfg.customer VALUES (86, 'Lazzy', 'Pankratz Place 49', 'lazzy.info', 882295.3903);
INSERT INTO wfg.customer VALUES (87, 'Browsecat', 'Thackeray Circle 28', 'browsecat.support', 408149.8052);
INSERT INTO wfg.customer VALUES (88, 'Aibox', 'Emmet Pass 59', 'aibox.xyz', 739861.4133);
INSERT INTO wfg.customer VALUES (89, 'Realbridge', 'Blaine Terrace 89', 'realbridge.support', 969848.8643);
INSERT INTO wfg.customer VALUES (90, 'Omba', 'Merchant Plaza 86', 'omba.app', 946286.9123);
INSERT INTO wfg.customer VALUES (91, 'Viva', 'Bluejay Center 68', 'viva.online', 417523.8863);
INSERT INTO wfg.customer VALUES (92, 'Brightdog', 'Village Green Terrace 22', 'brightdog.solutions', 43247.9614);
INSERT INTO wfg.customer VALUES (93, 'Voonder', 'Montana Plaza 99', 'voonder.net', 623918.664);
INSERT INTO wfg.customer VALUES (94, 'Gigabox', 'Buhler Drive 52', 'gigabox.biz', 116167.0683);
INSERT INTO wfg.customer VALUES (95, 'Myworks', 'Reinke Drive 25', 'myworks.global', 740472.9151);
INSERT INTO wfg.customer VALUES (96, 'Kare', 'Twin Pines Street 30', 'kare.info', 325714.032);
INSERT INTO wfg.customer VALUES (97, 'Wikivu', 'Meadow Valley Circle 21', 'wikivu.net', 516036.7636);
INSERT INTO wfg.customer VALUES (98, 'Gevee', 'Schurz Lane 72', 'gevee.xyz', 382974.8896);
INSERT INTO wfg.customer VALUES (99, 'Realfire', 'Lien Crossing 91', 'realfire.com', 726116.6874);
INSERT INTO wfg.customer VALUES (100, 'JumpXS', 'Morning Hill 8', 'jumpxs.me', 158957.2617);

-- INSERTS Account_type
INSERT INTO wfg.account_type(account_type_id,name,description)values(1,'Checking Account','Used for everyday transactions like paying bills, making purchases, and withdrawing cash.');
INSERT INTO wfg.account_type(account_type_id,name,description)values(2,'Savings Account','Designed to save money and earn interest over time.');
INSERT INTO wfg.account_type(account_type_id,name,description)values(3,'Business Account','Designed for business transactions and managing business finances.');

-- INSERTS Operation
INSERT INTO wfg.operation(operation_id,name,description)values(1,'Deposit','It consists of depositing money into the account');
INSERT INTO wfg.operation(operation_id,name,description)values(2,'Withdrawal','It consists of stealing money from an account');
INSERT INTO wfg.operation(operation_id,name,description)values(3,'Transfer','It consists of moving money from one account to another.');
