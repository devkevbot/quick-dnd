<template>
  <div class="profile-card">
    <v-container>
      <v-row wrap>
        <v-col>
          <v-card class="mb-2" elevation="2" outlined tile height="100%">
            <v-card-title class="display-1 primary--text">
              <v-icon class="mr-2" color="primary">mdi-account</v-icon>
              Account overview
            </v-card-title>

            <v-divider></v-divider>
            <v-spacer></v-spacer>

            <v-card-text class="primary--text headline"> Details </v-card-text>
            <v-card-text class="text--primary title">
              Username: {{ player.username }}
            </v-card-text>

            <v-spacer></v-spacer>

            <v-card-text class="text--primary title">
              Full name: {{ player.name }}</v-card-text
            >
            <v-divider></v-divider>
            <v-card-text class="primary--text headline"> Actions </v-card-text>

            <v-card-actions>
              <v-btn class="primary ml-2" @click="displayPasswordChangeForm">
                Change password
                <v-icon class="ml-2">mdi-lock-reset</v-icon>
              </v-btn>

              <ChangePasswordForm
                ref="confirmPasswordChange"
                @changePasswordFormEvent="requestChangePassword"
              />

              <v-btn class="error ml-2" @click="displayAccountDeletionPrompt">
                Delete account
                <v-icon class="ml-2">mdi-delete-forever</v-icon>
              </v-btn>

              <ConfirmDialog ref="confirmDelete" />
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import ChangePasswordForm from './ChangePasswordForm.vue';
import ConfirmDialog from '../ConfirmDialog.vue';

export default {
  name: 'ProfileCard',
  components: {
    ChangePasswordForm,
    ConfirmDialog,
  },
  data() {
    return {
      player: {
        username: '',
        name: '',
      },
    };
  },
  /**
   * Fetches all of the player's data when the component loads.
   */
  async created() {
    const loggedInPlayer = this.$store.getters['authentication/getPlayer'];
    const requestURI = `auth/player/${loggedInPlayer.username}`;
    const method = 'GET';

    await this.$http({
      url: requestURI,
      data: null,
      method,
    })
      .then((resp) => {
        this.player.username = resp.data.data.username;
        this.player.name = resp.data.data.name;
      })
      .catch(() => {
        /* TODO: Add error handling. */
      });
  },
  methods: {
    /**
     * Displays a password change form to the user.
     */
    async displayPasswordChangeForm() {
      await this.$refs.confirmPasswordChange.prompt();
    },
    /**
     * Sends a password change API request if the user has successfully
     * submitted the change password form.
     */
    async requestChangePassword({ success, newPassword, confirmation }) {
      if (!success) return;

      const requestURI = 'auth/player/me/password';
      const method = 'PUT';
      const data = {
        new_password: newPassword,
        confirmation,
      };

      await this.$http({
        url: requestURI,
        data,
        method,
      })
        .then(() => {
          /* TODO: Add success message. */
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
    /**
     * Displays a prompt to the user, asking them to confirm their
     * account deletion.
     */
    async displayAccountDeletionPrompt() {
      const title = 'Confirm deletion';
      const message = 'Delete your account? This action is irreversible.';
      if (await this.$refs.confirmDelete.prompt(title, message)) {
        await this.requestAccountDeletion();
      }
    },
    /**
     * Sends an HTTP request to delete the user's account.
     */
    async requestAccountDeletion() {
      const requestURI = 'auth/player/me';
      const method = 'DELETE';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then(() => {
          this.$router.push('/login');
          /* TODO: Add code to log the user out and clean up their
          authentication data. */
        })
        .catch(() => {
          /* TODO: Add error handling. */
        });
    },
  },
};
</script>
