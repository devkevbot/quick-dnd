<!-- CampaignForm.vue
This component is the view that allows the user to create a new
campaign. Assumes the user is authenticated.
-->

<template>
  <div class="campaign-form">
    <v-form>
      <v-container class="mb-1">
        <v-row class="ml-n3 mb-1">
          <p class="display-1 primary--text">Create a Campaign</p>
        </v-row>

        <v-row class="ml-n2 mt-2 mb-n1">
          <p class="subtitle-2 primary--text">Info</p>
        </v-row>

        <v-row class="mb-3">
          <v-col cols="12" md="6">
            <v-text-field
              v-model="name"
              :counter="50"
              label="Campaign Name"
              required
              :error-messages="nameErrors"
              @input="$v.name.$touch()"
              @blur="$v.name.$touch()"
            ></v-text-field>
          </v-col>

          <v-col cols="12" md="6">
            <v-text-field
              v-model="location"
              :counter="50"
              label="Location"
              required
              :error-messages="locationErrors"
              @input="$v.location.$touch()"
              @blur="$v.location.$touch()"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-textarea
            v-model="state"
            filled
            label="Current State"
            required
            :error-messages="stateErrors"
            @input="$v.state.$touch()"
            @blur="$v.state.$touch()"
          ></v-textarea>
        </v-row>

        <v-row class="ml-n2 mb-1">
          <p class="subtitle-2 primary--text">Campaign Characters</p>
        </v-row>

        <AddCharacterTable ref="addCharacterTable" />

        <v-row class="ml-0 mt-3">
          <v-btn @click="showAddCharacterDialog()">Add Character</v-btn>
        </v-row>

        <v-row class="ml-0 mt-10">
          <v-btn
            color="primary"
            @click="onClickCreate()"
          >Create</v-btn>
        </v-row>

        <AddCharacterDialog ref="addCharacterDialog" />
      </v-container>
    </v-form>
  </div>
</template>

<script>
import {
  required,
} from 'vuelidate/lib/validators';
import AddCharacterDialog from './AddCharacterDialog.vue';
import AddCharacterTable from './AddCharacterTable.vue';

export default {
  name: 'CampaignForm',
  components: {
    AddCharacterDialog,
    AddCharacterTable,
  },
  data() {
    return {
      name: '',
      location: '',
      state: '',
      idExists: false,
      tempCharacter: null,
    };
  },
  validations: {
    name: {
      required,
    },
    location: {
      required,
    },
    state: {
      required,
    },
  },
  methods: {
    /**
     * Displays prompt that asks user to enter a character ID
     */
    async onClickCreate() {
      this.$v.$touch();
      if (this.$v.$invalid) return;
      this.$v.$reset();
      this.requestCampaignCreation();
    },
    async showAddCharacterDialog() {
      if (await this.$refs.addCharacterDialog.prompt()) {
        const id = this.$refs.addCharacterDialog.getID;

        await this.fetchCharacterInfo(id);

        if (!this.idExists) return; // TODO: Notify user of invalid id

        this.$refs.addCharacterTable.addItem(
          id,
          this.tempCharacter.cName,
          this.tempCharacter.cOwner,
        );
      }
      this.$refs.addCharacterDialog.clearID();
      this.idExists = false;
      this.tempCharacter = null;
    },
    /**
     * Fetches character and owner names based on ID number, and places
     * the information in "tempCharacter"
     */
    async fetchCharacterInfo(id) {
      const integerID = parseInt(id, 10);
      const requestURI = `auth/character/${integerID}`;
      const method = 'GET';

      await this.$http({
        url: requestURI,
        data: null,
        method,
      })
        .then((resp) => {
          this.tempCharacter = {
            cName: resp.data.data.name,
            cOwner: resp.data.data.player_username,
          };
          this.idExists = true;
        })
        .catch(() => {
          /* TODO: Add error handling */
        });
    },
    /**
     * Sends HTTP requests to create a new campaign
     */
    async requestCampaignCreation() {
      const requestURI = 'auth/campaign';
      const data = {
        campaign: {
          name: this.name,
          current_location: this.location,
          state: this.state,
        },
        character_id: this.$refs.addCharacterTable.getIDs,
      };
      const method = 'POST';

      await this.$http({ url: requestURI, data, method })
        .then(() => {
          this.$router.push('/dashboard');
        })
        .catch(() => {
          /* TODO: Add error handling */
        });
    },
  },
  computed: {
    nameErrors() {
      const errors = [];
      if (!this.$v.name.$dirty) return errors;
      if (!this.$v.name.required) {
        errors.push('Name is required.');
      }
      return errors;
    },
    locationErrors() {
      const errors = [];
      if (!this.$v.location.$dirty) return errors;
      if (!this.$v.location.required) {
        errors.push('Location is required.');
      }
      return errors;
    },
    stateErrors() {
      const errors = [];
      if (!this.$v.state.$dirty) return errors;
      if (!this.$v.state.required) {
        errors.push('State is required.');
      }
      return errors;
    },
  },
};
</script>
