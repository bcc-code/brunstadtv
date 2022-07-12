import { Knex } from 'knex';

const episodes_access_view_sql = `
create or replace view episodes_roles(id, roles, roles_download, roles_earlyaccess) as
SELECT e.id,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups eu
                 WHERE eu.episodes_id = e.id), ARRAY []::character varying[]) AS roles,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups_download eu
                 WHERE eu.episodes_id = e.id), ARRAY []::character varying[]) AS roles_download,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups_earlyaccess eu
                 WHERE eu.episodes_id = e.id), ARRAY []::character varying[]) AS roles_earlyaccess
FROM episodes e;

create or replace view episodes_availability(id, published, available_from, available_to) as
SELECT e.id,
       e.status::text = 'published'::text AND se.status::text = 'published'::text AND
       s.status::text = 'published'::text                           AS published,
       COALESCE(GREATEST(e.available_from, se.available_from, s.available_from),
                '1800-01-01 00:00:00'::timestamp without time zone) AS available_from,
       COALESCE(LEAST(e.available_to, se.available_to, s.available_to),
                '3000-01-01 00:00:00'::timestamp without time zone) AS available_to
FROM episodes e
         LEFT JOIN seasons se ON e.season_id = se.id
         LEFT JOIN shows s ON se.show_id = s.id;
         
create or replace view episodes_access_view
            (id, published, available_from, available_to, usergroups, usergroups_downloads, usergroups_earlyaccess) as
SELECT e.id,
       ea.published,
       ea.available_from,
       ea.available_to,
       roles.roles             AS usergroups,
       roles.roles_download    AS usergroups_downloads,
       roles.roles_earlyaccess AS usergroups_earlyaccess
FROM episodes e
         LEFT JOIN episodes_availability ea ON ea.id = e.id
         LEFT JOIN episodes_roles roles ON roles.id = e.id;
`

const episodes_materialized_view_sql = `
create materialized view episodes_access as
SELECT episodes_access_view.id,
       episodes_access_view.published,
       episodes_access_view.available_from,
       episodes_access_view.available_to,
       episodes_access_view.usergroups,
       episodes_access_view.usergroups_downloads,
       episodes_access_view.usergroups_earlyaccess
FROM episodes_access_view;

create unique index episodes_access_idx
    on episodes_access (id);
`

const seasons_access_view_sql = `
create or replace view seasons_roles(id, roles, roles_download, roles_earlyaccess) as
SELECT se.id,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups eu
                 WHERE (eu.episodes_id IN (SELECT e.id
                                           FROM episodes e
                                           WHERE e.season_id = se.id))), ARRAY []::character varying[]) AS roles,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups_download eu
                 WHERE (eu.episodes_id IN (SELECT e.id
                                           FROM episodes e
                                           WHERE e.season_id = se.id))),
                ARRAY []::character varying[])                                                          AS roles_download,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups_earlyaccess eu
                 WHERE (eu.episodes_id IN (SELECT e.id
                                           FROM episodes e
                                           WHERE e.season_id = se.id))),
                ARRAY []::character varying[])                                                          AS roles_earlyaccess
FROM seasons se;

create or replace view seasons_availability(id, published, available_from, available_to) as
SELECT se.id,
       se.status::text = 'published'::text AND s.status::text = 'published'::text                                  AS published,
       COALESCE(GREATEST(se.available_from, s.available_from),
                '1800-01-01 00:00:00'::timestamp without time zone)                                                AS available_from,
       COALESCE(LEAST(se.available_to, s.available_to),
                '3000-01-01 00:00:00'::timestamp without time zone)                                                AS available_to
FROM seasons se
         LEFT JOIN shows s ON se.show_id = s.id;
         
create or replace view seasons_access_view
	(id, published, available_from, available_to, usergroups, usergroups_downloads, usergroups_earlyaccess) as
SELECT se.id,
       a.published,
       a.available_from,
       a.available_to,
       r.roles             AS usergroups,
       r.roles_download    AS usergroups_downloads,
       r.roles_earlyaccess AS usergroups_earlyaccess
FROM seasons se
         LEFT JOIN seasons_availability a ON a.id = se.id
         LEFT JOIN seasons_roles r ON r.id = se.id;
`;

const seasons_materialized_view_sql = `
create materialized view seasons_access as
SELECT seasons_access_view.id,
       seasons_access_view.published,
       seasons_access_view.available_from,
       seasons_access_view.available_to,
       seasons_access_view.usergroups,
       seasons_access_view.usergroups_downloads,
       seasons_access_view.usergroups_earlyaccess
FROM seasons_access_view;

create unique index seasons_access_idx
    on seasons_access (id);
`

const shows_access_view_sql = `
create view shows_roles(id, roles, roles_download, roles_earlyaccess) as
SELECT sh.id,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups eu
                 WHERE (eu.episodes_id IN (SELECT e.id
                                           FROM episodes e
                                           WHERE (e.season_id IN (SELECT se.id
                                                                  FROM seasons se
                                                                  WHERE se.show_id = sh.id))))),
                ARRAY []::character varying[]) AS roles,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups_download eu
                 WHERE (eu.episodes_id IN (SELECT e.id
                                           FROM episodes e
                                           WHERE (e.season_id IN (SELECT se.id
                                                                  FROM seasons se
                                                                  WHERE se.show_id = sh.id))))),
                ARRAY []::character varying[]) AS roles_download,
       COALESCE((SELECT array_agg(DISTINCT eu.usergroups_code) AS code
                 FROM episodes_usergroups_earlyaccess eu
                 WHERE (eu.episodes_id IN (SELECT e.id
                                           FROM episodes e
                                           WHERE (e.season_id IN (SELECT se.id
                                                                  FROM seasons se
                                                                  WHERE se.show_id = sh.id))))),
                ARRAY []::character varying[]) AS roles_earlyaccess
FROM shows sh;

create view shows_availability(id, published, available_from, available_to) as
SELECT sh.id,
       sh.status::text = 'published'::text                                             AS published,
       COALESCE(sh.available_from, '1800-01-01 00:00:00'::timestamp without time zone) AS available_from,
       COALESCE(sh.available_to, '3000-01-01 00:00:00'::timestamp without time zone)   AS available_to
FROM shows sh;

create view shows_access_view
            (id, published, available_from, available_to, usergroups, usergroups_downloads, usergroups_earlyaccess) as
SELECT sh.id,
       a.published,
       a.available_from,
       a.available_to,
       r.roles             AS usergroups,
       r.roles_download    AS usergroups_downloads,
       r.roles_earlyaccess AS usergroups_earlyaccess
FROM shows sh
         LEFT JOIN shows_availability a ON a.id = sh.id
         LEFT JOIN shows_roles r ON r.id = sh.id;
`;

const shows_materialized_view_sql = `
create materialized view shows_access as
SELECT shows_access_view.id,
       shows_access_view.published,
       shows_access_view.available_from,
       shows_access_view.available_to,
       shows_access_view.usergroups,
       shows_access_view.usergroups_downloads,
       shows_access_view.usergroups_earlyaccess
FROM shows_access_view;

create unique index shows_access_idx
    on shows_access (id);
`

module.exports = {
	async up(k : Knex) {
		await k.raw(episodes_access_view_sql);
		//TODO: replace existing materialized view?
		//await k.raw(episodes_materialized_view_sql);

		await k.raw(seasons_access_view_sql);
		await k.raw(seasons_materialized_view_sql);
		await k.raw(shows_access_view_sql);
		await k.raw(shows_materialized_view_sql);
	},

	async down(k : Knex) {
		//await k.raw(`DROP MATERIALIZED VIEW IF EXISTS seasons_access`);
		await k.raw(`DROP VIEW IF EXISTS episodes_access_view`);
		await k.raw(`DROP VIEW IF EXISTS episodes_roles`);
		await k.raw(`DROP VIEW IF EXISTS episodes_availability`);
		await k.raw(`DROP MATERIALIZED VIEW IF EXISTS seasons_access`);
		await k.raw(`DROP VIEW IF EXISTS seasons_access_view`);
		await k.raw(`DROP VIEW IF EXISTS seasons_roles`);
		await k.raw(`DROP VIEW IF EXISTS seasons_availability`);
		await k.raw(`DROP MATERIALIZED VIEW IF EXISTS shows_access`);
		await k.raw(`DROP VIEW IF EXISTS shows_access_view`);
		await k.raw(`DROP VIEW IF EXISTS shows_roles`);
		await k.raw(`DROP VIEW IF EXISTS shows_availability`);
	}
}
